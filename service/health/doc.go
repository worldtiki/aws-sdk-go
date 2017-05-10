// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package health provides the client and types for making API
// requests to AWS Health APIs and Notifications.
//
// The AWS Health API provides programmatic access to the AWS Health information
// that is presented in the AWS Personal Health Dashboard (https://phd.aws.amazon.com/phd/home#/).
// You can get information about events that affect your AWS resources:
//
//    * DescribeEvents: Summary information about events.
//
//    * DescribeEventDetails: Detailed information about one or more events.
//
//    * DescribeAffectedEntities: Information about AWS resources that are affected
//    by one or more events.
//
// In addition, these operations provide information about event types and summary
// counts of events or affected entities:
//
//    * DescribeEventTypes: Information about the kinds of events that AWS Health
//    tracks.
//
//    * DescribeEventAggregates: A count of the number of events that meet specified
//    criteria.
//
//    * DescribeEntityAggregates: A count of the number of affected entities
//    that meet specified criteria.
//
// The Health API requires a Business or Enterprise support plan from AWS Support
// (http://aws.amazon.com/premiumsupport/). Calling the Health API from an account
// that does not have a Business or Enterprise support plan causes a SubscriptionRequiredException.
//
// For authentication of requests, AWS Health uses the Signature Version 4 Signing
// Process (http://docs.aws.amazon.com/general/latest/gr/signature-version-4.html).
//
// See the AWS Health User Guide (http://docs.aws.amazon.com/health/latest/ug/what-is-aws-health.html)
// for information about how to use the API.
//
// Service Endpoint
//
// The HTTP endpoint for the AWS Health API is:
//
//    * https://health.us-east-1.amazonaws.com
//
// See https://docs.aws.amazon.com/goto/WebAPI/health-2016-08-04 for more information on this service.
//
// See health package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/health/
//
// Using the Client
//
// To use the client for AWS Health APIs and Notifications you will first need
// to create a new instance of it.
//
// When creating a client for an AWS service you'll first need to have a Session
// already created. The Session provides configuration that can be shared
// between multiple service clients. Additional configuration can be applied to
// the Session and service's client when they are constructed. The aws package's
// Config type contains several fields such as Region for the AWS Region the
// client should make API requests too. The optional Config value can be provided
// as the variadic argument for Sessions and client creation.
//
// Once the service's client is created you can use it to make API requests the
// AWS service. These clients are safe to use concurrently.
//
//   // Create a session to share configuration, and load external configuration.
//   sess := session.Must(session.NewSession())
//
//   // Create the service's client with the session.
//   svc := health.New(sess)
//
// See the SDK's documentation for more information on how to use service clients.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws package's Config type for more information on configuration options.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the AWS Health APIs and Notifications client Health for more
// information on creating the service's client.
// https://docs.aws.amazon.com/sdk-for-go/api/service/health/#New
//
// Once the client is created you can make an API request to the service.
// Each API method takes a input parameter, and returns the service response
// and an error.
//
// The API method will document which error codes the service can be returned
// by the operation if the service models the API operation's errors. These
// errors will also be available as const strings prefixed with "ErrCode".
//
//   result, err := svc.DescribeAffectedEntities(params)
//   if err != nil {
//       // Cast err to awserr.Error to handle specific error codes.
//       aerr, ok := err.(awserr.Error)
//       if ok && aerr.Code() == <error code to check for> {
//           // Specific error code handling
//       }
//       return err
//   }
//
//   fmt.Println("DescribeAffectedEntities result:")
//   fmt.Println(result)
//
// Using the Client with Context
//
// The service's client also provides methods to make API requests with a Context
// value. This allows you to control the timeout, and cancellation of pending
// requests. These methods also take request Option as variadic parameter to apply
// additional configuration to the API request.
//
//   ctx := context.Background()
//
//   result, err := svc.DescribeAffectedEntitiesWithContext(ctx, params)
//
// See the request package documentation for more information on using Context pattern
// with the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/request/
package health