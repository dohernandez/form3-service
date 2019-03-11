Feature: Read payment
  As a user I want to see a payment
  So that I can check all the payment's attributes

  Background:
    Given that the following payment state(s) are stored in the table "events_transaction_stream"

      | event_id                             | event_name                     | metadata                                                                                                         | created_at | payload                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
      | 8b3f0880-800c-40cf-9cc6-2d53be233c3f | transaction_payment_created_v0 | {"_aggregate_id": "71aa6f04-ede9-46f4-a63d-373c3c206fc1", "_aggregate_type": "Payment", "_aggregate_version": 1} | now        | {"id":"71aa6f04-ede9-46f4-a63d-373c3c206fc1","organisation_id":"743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb","attributes":{"amount":"100.21","beneficiary_party":{"name":"Wilfred Jeremiah Owens","address":"1 The Beneficiary Localtown SE2","bank_id":"403000","bank_id_code":"GBDSC","account_number":"31926819","account_name":"W Owens","account_number_code":"BBAN"},"charges_information":{"bearer_code":"SHAR","sender_charges":[{"amount":"5","currency":"GBP"},{"amount":"10","currency":"USD"}],"receiver_charges_amount":"1","receiver_charges_currency":"USD"},"currency":"GBP","debtor_party":{"name":"Emelia Jane Brown","address":"10 Debtor Crescent Sourcetown NE1","bank_id":"203301","bank_id_code":"GBDSC","account_number":"GB29XABC10161234567801","account_name":"EJ Brown Black","account_number_code":"IBAN"},"end_to_end_reference":"Wil piano Jan","fx":{"contract_reference":"FX123","exchange_rate":"2","original_amount":"200.42","original_currency":"USD"},"numeric_reference":"1002001","processing_date":"2017-01-18","reference":"Payment for Em's piano lessons","sponsor_party":{"bank_id":"123123","bank_id_code":"GBDSC","account_number":"56781234"},"payment_id":"123456789012345678","payment_purpose":"Paying for goods/services","payment_scheme":"FPS","payment_type":"Credit","scheme_payment_type":"ImmediatePayment","scheme_payment_sub_type":"InternetBanking"}} |
      | 5c6a6ae0-d005-4123-bfb0-4758594ae3b8 | transaction_payment_created_v0 | {"_aggregate_id": "5f6147a5-b2fd-4128-afe6-9b1b0e7534ee", "_aggregate_type": "Payment", "_aggregate_version": 1} | now        | {"id":"5f6147a5-b2fd-4128-afe6-9b1b0e7534ee","organisation_id":"743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb","attributes":{"amount":"100.21","beneficiary_party":{"name": "Wilfred Jeremiah Connor", "address": "1 The Beneficiary Localtown SE2", "bank_id": "403000", "account_name": "W Connor", "bank_id_code": "GBDSC", "account_number": "31926812", "account_number_code": "BBAN"},"charges_information":{"bearer_code":"SHAR","sender_charges":[{"amount":"5","currency":"GBP"},{"amount":"10","currency":"USD"}],"receiver_charges_amount":"1","receiver_charges_currency":"USD"},"currency":"GBP","debtor_party":{"name":"Emelia Jane Brown","address":"10 Debtor Crescent Sourcetown NE1","bank_id":"203301","bank_id_code":"GBDSC","account_number":"GB29XABC10161234567801","account_name":"EJ Brown Black","account_number_code":"IBAN"},"end_to_end_reference":"Wil piano Jan","fx":{"contract_reference":"FX123","exchange_rate":"2","original_amount":"200.42","original_currency":"USD"},"numeric_reference":"1002001","processing_date":"2017-01-18","reference":"Payment for Em's piano lessons","sponsor_party":{"bank_id":"123123","bank_id_code":"GBDSC","account_number":"56781234"},"payment_id":"123456789012345678","payment_purpose":"Paying for goods/services","payment_scheme":"FPS","payment_type":"Credit","scheme_payment_type":"ImmediatePayment","scheme_payment_sub_type":"InternetBanking"}} |

    And that the following payment(s) are stored in the table "transaction_payment"

      | id                                   | version | organisation_id                      | attributes                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
      | 71aa6f04-ede9-46f4-a63d-373c3c206fc1 | 0       | 743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb | {"fx": {"exchange_rate": "2", "original_amount": "200.42", "original_currency": "USD", "contract_reference": "FX123"}, "amount": "100.21", "currency": "GBP", "reference": "Payment for Em's piano lessons", "payment_id": "123456789012345678", "debtor_party": {"name": "Emelia Jane Brown", "address": "10 Debtor Crescent Sourcetown NE1", "bank_id": "203301", "account_name": "EJ Brown Black", "bank_id_code": "GBDSC", "account_number": "GB29XABC10161234567801", "account_number_code": "IBAN"}, "payment_type": "Credit", "sponsor_party": {"bank_id": "123123", "bank_id_code": "GBDSC", "account_number": "56781234"}, "payment_scheme": "FPS", "payment_purpose": "Paying for goods/services", "processing_date": "2017-01-18", "beneficiary_party": {"name": "Wilfred Jeremiah Owens", "address": "1 The Beneficiary Localtown SE2", "bank_id": "403000", "account_name": "W Owens", "bank_id_code": "GBDSC", "account_number": "31926819", "account_number_code": "BBAN"}, "numeric_reference": "1002001", "charges_information": {"bearer_code": "SHAR", "sender_charges": [{"amount": "5", "currency": "GBP"}, {"amount": "10", "currency": "USD"}], "receiver_charges_amount": "1", "receiver_charges_currency": "USD"}, "scheme_payment_type": "ImmediatePayment", "end_to_end_reference": "Wil piano Jan", "scheme_payment_sub_type": "InternetBanking"}   |
      | 5f6147a5-b2fd-4128-afe6-9b1b0e7534ee | 0       | 743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb | {"fx": {"exchange_rate": "2", "original_amount": "200.42", "original_currency": "USD", "contract_reference": "FX123"}, "amount": "100.21", "currency": "GBP", "reference": "Payment for Em's piano lessons", "payment_id": "123456789012345678", "debtor_party": {"name": "Emelia Jane Brown", "address": "10 Debtor Crescent Sourcetown NE1", "bank_id": "203301", "account_name": "EJ Brown Black", "bank_id_code": "GBDSC", "account_number": "GB29XABC10161234567801", "account_number_code": "IBAN"}, "payment_type": "Credit", "sponsor_party": {"bank_id": "123123", "bank_id_code": "GBDSC", "account_number": "56781234"}, "payment_scheme": "FPS", "payment_purpose": "Paying for goods/services", "processing_date": "2017-01-18", "beneficiary_party": {"name": "Wilfred Jeremiah Connor", "address": "1 The Beneficiary Localtown SE2", "bank_id": "403000", "account_name": "W Connor", "bank_id_code": "GBDSC", "account_number": "31926812", "account_number_code": "BBAN"}, "numeric_reference": "1002001", "charges_information": {"bearer_code": "SHAR", "sender_charges": [{"amount": "5", "currency": "GBP"}, {"amount": "10", "currency": "USD"}], "receiver_charges_amount": "1", "receiver_charges_currency": "USD"}, "scheme_payment_type": "ImmediatePayment", "end_to_end_reference": "Wil piano Jan", "scheme_payment_sub_type": "InternetBanking"} |

  Scenario: Payment should be deleted
    When I request REST endpoint with method "GET" and path "/v1/transaction/payments"

    Then I should have an OK response

    And I should have a response with following JSON body
    """
    {
        "data": [
            {
                "id": "71aa6f04-ede9-46f4-a63d-373c3c206fc1",
                "version": 0,
                "organisation_id": "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",
                "attributes": {
                    "amount": "100.21",
                    "beneficiary_party": {
                        "name": "Wilfred Jeremiah Owens",
                        "address": "1 The Beneficiary Localtown SE2",
                        "bank_id": "403000",
                        "bank_id_code": "GBDSC",
                        "account_number": "31926819",
                        "account_name": "W Owens",
                        "account_number_code": "BBAN"
                    },
                    "charges_information": {
                        "bearer_code": "SHAR",
                        "sender_charges": [
                            {
                                "amount": "5",
                                "currency": "GBP"
                            },
                            {
                                "amount": "10",
                                "currency": "USD"
                            }
                        ],
                        "receiver_charges_amount": "1",
                        "receiver_charges_currency": "USD"
                    },
                    "currency": "GBP",
                    "debtor_party": {
                        "name": "Emelia Jane Brown",
                        "address": "10 Debtor Crescent Sourcetown NE1",
                        "bank_id": "203301",
                        "bank_id_code": "GBDSC",
                        "account_number": "GB29XABC10161234567801",
                        "account_name": "EJ Brown Black",
                        "account_number_code": "IBAN"
                    },
                    "end_to_end_reference": "Wil piano Jan",
                    "fx": {
                        "contract_reference": "FX123",
                        "exchange_rate": "2",
                        "original_amount": "200.42",
                        "original_currency": "USD"
                    },
                    "numeric_reference": "1002001",
                    "processing_date": "2017-01-18",
                    "reference": "Payment for Em's piano lessons",
                    "sponsor_party": {
                        "bank_id": "123123",
                        "bank_id_code": "GBDSC",
                        "account_number": "56781234"
                    },
                    "payment_id": "123456789012345678",
                    "payment_purpose": "Paying for goods/services",
                    "payment_scheme": "FPS",
                    "payment_type": "Credit",
                    "scheme_payment_type": "ImmediatePayment",
                    "scheme_payment_sub_type": "InternetBanking"
                },
                "type": "Payment"
            },
            {
                "id": "5f6147a5-b2fd-4128-afe6-9b1b0e7534ee",
                "version": 0,
                "organisation_id": "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",
                "attributes": {
                    "amount": "100.21",
                    "beneficiary_party": {
                        "name": "Wilfred Jeremiah Connor",
                        "address": "1 The Beneficiary Localtown SE2",
                        "bank_id": "403000",
                        "bank_id_code": "GBDSC",
                        "account_number": "31926812",
                        "account_name": "W Connor",
                        "account_number_code": "BBAN"
                    },
                    "charges_information": {
                        "bearer_code": "SHAR",
                        "sender_charges": [
                            {
                                "amount": "5",
                                "currency": "GBP"
                            },
                            {
                                "amount": "10",
                                "currency": "USD"
                            }
                        ],
                        "receiver_charges_amount": "1",
                        "receiver_charges_currency": "USD"
                    },
                    "currency": "GBP",
                    "debtor_party": {
                        "name": "Emelia Jane Brown",
                        "address": "10 Debtor Crescent Sourcetown NE1",
                        "bank_id": "203301",
                        "bank_id_code": "GBDSC",
                        "account_number": "GB29XABC10161234567801",
                        "account_name": "EJ Brown Black",
                        "account_number_code": "IBAN"
                    },
                    "end_to_end_reference": "Wil piano Jan",
                    "fx": {
                        "contract_reference": "FX123",
                        "exchange_rate": "2",
                        "original_amount": "200.42",
                        "original_currency": "USD"
                    },
                    "numeric_reference": "1002001",
                    "processing_date": "2017-01-18",
                    "reference": "Payment for Em's piano lessons",
                    "sponsor_party": {
                        "bank_id": "123123",
                        "bank_id_code": "GBDSC",
                        "account_number": "56781234"
                    },
                    "payment_id": "123456789012345678",
                    "payment_purpose": "Paying for goods/services",
                    "payment_scheme": "FPS",
                    "payment_type": "Credit",
                    "scheme_payment_type": "ImmediatePayment",
                    "scheme_payment_sub_type": "InternetBanking"
                },
                "type": "Payment"
            }
        ],
        "links": {
            "self": "/v1/transaction/payments"
        }
    }
    """