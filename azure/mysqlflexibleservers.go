package azure

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/mysql/mgmt/mysqlflexibleservers"
	"github.com/gruntwork-io/terratest/modules/testing"
	"github.com/stretchr/testify/require"
)

// GetMYSQLServerClientE is a helper function that will setup a mysql server client.
// TODO: remove in next version
func CreateMySQLFlexibleServerClientE(subscriptionID string) (*mysqlflexibleservers.ServersClient, error) {
	// Validate Azure subscription ID
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return nil, err
	}

	// Create a mysql server client
	mysqlClient := mysqlflexibleservers.NewServersClient(subscriptionID)

	// Create an authorizer
	authorizer, err := NewAuthorizer()
	if err != nil {
		return nil, err
	}

	// Attach authorizer to the client
	mysqlClient.Authorizer = *authorizer

	return &mysqlClient, nil
}

// GetMYSQLServer is a helper function that gets the server.
// This function would fail the test if there is an error.
func GetMySQLFlexibleServer(t testing.TestingT, resGroupName string, serverName string, subscriptionID string) *mysqlflexibleservers.Server {
	mysqlServer, err := GetMySQLFlexibleServerE(t, subscriptionID, resGroupName, serverName)
	require.NoError(t, err)

	return mysqlServer
}

// GetMYSQLServerE is a helper function that gets the server.
func GetMySQLFlexibleServerE(t testing.TestingT, subscriptionID string, resGroupName string, serverName string) (*mysqlflexibleservers.Server, error) {
	// Create a mySQl Server client
	mysqlClient, err := CreateMySQLFlexibleServerClientE(subscriptionID)
	if err != nil {
		return nil, err
	}

	// Get the corresponding server client
	mysqlServer, err := mysqlClient.Get(context.Background(), resGroupName, serverName)
	if err != nil {
		return nil, err
	}

	//Return server
	return &mysqlServer, nil
}

// GetMYSQLDBClientE is a helper function that will setup a mysql DB client.
func GetMySQLFlexibleDBClientE(subscriptionID string) (*mysqlflexibleservers.DatabasesClient, error) {
	// Validate Azure subscription ID
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return nil, err
	}

	// Create a mysql db client
	mysqlDBClient := mysqlflexibleservers.NewDatabasesClient(subscriptionID)

	// Create an authorizer
	authorizer, err := NewAuthorizer()
	if err != nil {
		return nil, err
	}

	// Attach authorizer to the client
	mysqlDBClient.Authorizer = *authorizer

	return &mysqlDBClient, nil
}

// GetMYSQLDB is a helper function that gets the database.
// This function would fail the test if there is an error.
func GetMySQLFlexibleDB(t testing.TestingT, resGroupName string, serverName string, dbName string, subscriptionID string) *mysqlflexibleservers.Database {
	database, err := GetMySQLFlexibleDBE(t, subscriptionID, resGroupName, serverName, dbName)
	require.NoError(t, err)

	return database
}

// GetMYSQLDBE is a helper function that gets the database.
func GetMySQLFlexibleDBE(t testing.TestingT, subscriptionID string, resGroupName string, serverName string, dbName string) (*mysqlflexibleservers.Database, error) {
	// Create a mySQl db client
	mysqldbClient, err := GetMySQLFlexibleDBClientE(subscriptionID)
	if err != nil {
		return nil, err
	}

	// Get the corresponding db client
	mysqlDb, err := mysqldbClient.Get(context.Background(), resGroupName, serverName, dbName)
	if err != nil {
		return nil, err
	}

	//Return DB
	return &mysqlDb, nil
}

// ListMySQLDB is a helper function that gets all databases per server.
func ListMySQLFlexibleDB(t testing.TestingT, resGroupName string, serverName string, subscriptionID string) []mysqlflexibleservers.Database {
	dblist, err := ListMySQLFlexibleDBE(t, subscriptionID, resGroupName, serverName)
	require.NoError(t, err)

	return dblist
}

// ListMySQLDBE is a helper function that gets all databases per server.
func ListMySQLFlexibleDBE(t testing.TestingT, subscriptionID string, resGroupName string, serverName string) ([]mysqlflexibleservers.Database, error) {
	// Create a mySQl db client
	mysqldbClient, err := GetMySQLFlexibleDBClientE(subscriptionID)
	if err != nil {
		return nil, err
	}

	// Get the corresponding db client
	mysqlDbs, err := mysqldbClient.ListByServer(context.Background(), resGroupName, serverName)
	if err != nil {
		return nil, err
	}

	//Return DB lists
	return *mysqlDbs.Response().Value, nil
}
