package app

import (
	"context"
	"runtime"
	"time"

	"github.com/dohernandez/form3-service/internal/domain/transaction"
	"github.com/dohernandez/form3-service/internal/platform/event/store"
	message "github.com/dohernandez/form3-service/internal/platform/projection/handler/message"
	"github.com/dohernandez/form3-service/internal/platform/storage"
	"github.com/dohernandez/form3-service/pkg/app"
	"github.com/dohernandez/form3-service/pkg/event"
	"github.com/dohernandez/form3-service/pkg/log"
	"github.com/dohernandez/form3-service/pkg/projection"
	"github.com/hellofresh/goengine"
	"github.com/hellofresh/goengine/aggregate"
	driverSQL "github.com/hellofresh/goengine/driver/sql"
	"github.com/hellofresh/goengine/extension/pq"
	"github.com/hellofresh/goengine/strategy/json"
	"github.com/hellofresh/goengine/strategy/json/sql/postgres"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // enable postgres driver
	"github.com/sirupsen/logrus"
)

const (
	transactionStream = "transaction_stream"

	transactionProjectionsTable = "transaction_projections"
	transactionPaymentTable     = "transaction_payment"

	transactionPaymentCreated۰v0 = "transaction_payment_created_v0"
)

// NewAppContainer initializes application container
func NewAppContainer(cfg Config) (*Container, error) {
	// Create base container
	bc, err := app.NewAppContainer(cfg.Config)
	if err != nil {
		return nil, err
	}

	// Create application container
	c := newContainer(cfg, bc)

	// Init db
	db, err := sqlx.Open("postgres", cfg.DatabaseDSN)
	if err != nil {
		return nil, err
	}
	c.WithDB(db)

	// Create goengine postgres json SingleStreamManager
	manager, err := postgres.NewSingleStreamManager(db.DB, goengine.NopLogger)
	if err != nil {
		return nil, err
	}

	// Set the name of an event stream
	streamName := goengine.StreamName(transactionStream)

	paymentRepo, err := newPaymentRepository(manager, streamName)
	if err != nil {
		return nil, err
	}

	eventStore := event.NewStore(paymentRepo)
	c.WithPaymentEventStore(eventStore)

	if err := listenPaymentProjection(db, manager, streamName, cfg, c.Logger()); err != nil {
		return nil, err
	}

	return c, nil
}

// newPaymentRepository instantiates a new Payment AggregateRepository
func newPaymentRepository(
	manager *postgres.SingleStreamManager,
	streamName goengine.StreamName,
) (*aggregate.Repository, error) {
	// Register your events so that can be properly loaded from the event eventStore
	if err := manager.RegisterPayloads(map[string]json.PayloadInitiator{
		transactionPaymentCreated۰v0: func() interface{} {
			return transaction.PaymentCreated۰v0{}
		},
	}); err != nil {
		return nil, err
	}

	// Creates a new event store instance
	eventStore, err := manager.NewEventStore()
	if err != nil {
		return nil, err
	}

	// Creates a new aggregate.Type instance to be used to reconstitute the transaction.Payment version 0
	aggregatePaymentType, err := store.NewTransactionPaymentType۰v0()
	if err != nil {
		return nil, err
	}

	repo, err := aggregate.NewRepository(eventStore, streamName, aggregatePaymentType)
	if err != nil {
		return nil, err
	}

	return repo, nil
}

// listenPaymentProjection executes the payment projection and listens to any changes to the event store
func listenPaymentProjection(
	db *sqlx.DB,
	manager *postgres.SingleStreamManager,
	streamName goengine.StreamName,
	cfg Config,
	logger *logrus.Logger,
) error {
	ctx := log.ToContext(context.Background(), logger)

	paymentStorage := storage.NewPaymentStorage(db, transactionPaymentTable)

	projection := projection.NewProjection(transactionPaymentTable, streamName, func() interface{} {
		return store.PaymentState()
	})

	projection.RegisterMessageHandlers(map[string]goengine.MessageHandler{
		transactionPaymentCreated۰v0: message.TransactionPaymentCreatedHandler۰v0(
			paymentStorage,
		),
	})

	projector, err := manager.NewStreamProjector(
		transactionProjectionsTable,
		projection,
		func(err error, notification *driverSQL.ProjectionNotification) driverSQL.ProjectionErrorAction {
			return driverSQL.ProjectionFail
		},
	)
	if err != nil {
		return err
	}

	listener, err := pq.NewListener(
		cfg.DatabaseDSN,
		string(projection.FromStream()),
		time.Millisecond,
		time.Second,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		if err := projector.RunAndListen(ctx, listener); err != nil {
			logger.Fatalf("Failed executes the projection and listens %s", err.Error())
		}
	}()
	runtime.Gosched()

	return nil
}
