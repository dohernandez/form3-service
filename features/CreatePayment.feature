Feature: Create payment
  As a user I want to create a payment
  So that I can see later all my payments

  Scenario: Payment should be created
    When I request REST endpoint with method "POST" and path "/v1/transaction/payments" and body
    """
    {
	  "organisation_id": "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",
	  "attributes": {
        "amount": "100.21",
        "beneficiary_party": {
          "account_name": "W Owens",
          "account_number": "31926819",
          "account_number_code": "BBAN",
          "account_type": 0,
          "address": "1 The Beneficiary Localtown SE2",
          "bank_id": "403000",
          "bank_id_code": "GBDSC",
          "name": "Wilfred Jeremiah Owens"
        },
        "charges_information": {
          "bearer_code": "SHAR",
          "sender_charges": [
            {
              "amount": "5.00",
              "currency": "GBP"
            },
            {
              "amount": "10.00",
              "currency": "USD"
            }
          ],
          "receiver_charges_amount": "1.00",
          "receiver_charges_currency": "USD"
        },
        "currency": "GBP",
        "debtor_party": {
          "account_name": "EJ Brown Black",
          "account_number": "GB29XABC10161234567801",
          "account_number_code": "IBAN",
          "address": "10 Debtor Crescent Sourcetown NE1",
          "bank_id": "203301",
          "bank_id_code": "GBDSC",
          "name": "Emelia Jane Brown"
        },
        "end_to_end_reference": "Wil piano Jan",
        "fx": {
          "contract_reference": "FX123",
          "exchange_rate": "2.00000",
          "original_amount": "200.42",
          "original_currency": "USD"
        },
        "numeric_reference": "1002001",
        "payment_id": "123456789012345678",
        "payment_purpose": "Paying for goods/services",
        "payment_scheme": "FPS",
        "payment_type": "Credit",
        "processing_date": "2017-01-18",
        "reference": "Payment for Em's piano lessons",
        "scheme_payment_sub_type": "InternetBanking",
        "scheme_payment_type": "ImmediatePayment",
        "sponsor_party": {
          "account_number": "56781234",
          "bank_id": "123123",
          "bank_id_code": "GBDSC"
        }
      }
    }
    """

    Then I should have a created response
    And the following payment(s) should be stored in the table "transaction_payment"

      | ID  | version | organisation_id                      | attributes                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
      | $p1 | 0       | 743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb | {"fx": {"exchange_rate": "2", "original_amount": "200.42", "original_currency": "USD", "contract_reference": "FX123"}, "amount": "100.21", "currency": "GBP", "reference": "Payment for Em's piano lessons", "payment_id": "123456789012345678", "debtor_party": {"name": "Emelia Jane Brown", "address": "10 Debtor Crescent Sourcetown NE1", "bank_id": "203301", "account_name": "EJ Brown Black", "bank_id_code": "GBDSC", "account_number": "GB29XABC10161234567801", "account_number_code": "IBAN"}, "payment_type": "Credit", "sponsor_party": {"bank_id": "123123", "bank_id_code": "GBDSC", "account_number": "56781234"}, "payment_scheme": "FPS", "payment_purpose": "Paying for goods/services", "processing_date": "2017-01-18", "beneficiary_party": {"name": "Wilfred Jeremiah Owens", "address": "1 The Beneficiary Localtown SE2", "bank_id": "403000", "account_name": "W Owens", "bank_id_code": "GBDSC", "account_number": "31926819", "account_number_code": "BBAN"}, "numeric_reference": "1002001", "charges_information": {"bearer_code": "SHAR", "sender_charges": [{"amount": "5", "currency": "GBP"}, {"amount": "10", "currency": "USD"}], "receiver_charges_amount": "1", "receiver_charges_currency": "USD"}, "scheme_payment_type": "ImmediatePayment", "end_to_end_reference": "Wil piano Jan", "scheme_payment_sub_type": "InternetBanking"} |

    And the following payment state should be stored in the table "events_transaction_stream"

      | event_name                     | METADATA                                                                        |
      | transaction_payment_created_v0 | {"_aggregate_id": "$p1", "_aggregate_type": "Payment", "_aggregate_version": 1} |

  Scenario: Payment should not be created attribute missing
    When I request REST endpoint with method "POST" and path "/v1/transaction/payments" and body
    """
    {
	  "organisation_id": "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb"
    }
    """

    Then I should have a precondition failed response