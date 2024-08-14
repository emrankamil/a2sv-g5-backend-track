package usecases

import (
	"context"
	"log"
	"testing"
	domain "testing_task-manager_api/Domain"
	repositories "testing_task-manager_api/Repositories"
	"testing_task-manager_api/config"
	"testing_task-manager_api/mocks"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userUsecaseSuite struct{
	suite.Suite
	client    *mongo.Client
	db        *mongo.Database
	repository *mocks.UserRepository
	usecase domain.UserUsecase
}

func (suite *userUsecaseSuite) SetupSuite(){
	configs := config.GetConfig()
	client, db := config.ConnectDB(configs)

	repository := new(mocks.UserRepository)
	usecase := NewUserUsecase(repository, 10)

	suite.client = client
	suite.db = db
	suite.repository = repository
	suite.usecase = usecase
}

func (suite *userUsecaseSuite) TearDownSuite() {
	// Drop the test database after all tests are run
	if suite.db != nil {
		suite.db.Drop(context.TODO())
	}

	// Close the MongoDB connection
	config.CloseMongoDBConnection(suite.client)
}

func (suite *userUsecaseSuite) TearDownTest() {
	// List of collections you might want to clear after each test
	collections := []string{"users"}

	for _, collection := range collections {
		_, err := suite.db.Collection(collection).DeleteMany(context.TODO(), bson.D{})
		if err != nil {
			log.Fatalf("Failed to clear collection %s: %v", collection, err)
		}
	}
}
func (suite *userUsecaseSuite) SetupTest() {
	// Initialize the repository with the real database
	repository := repositories.NewUserRepository(suite.db, "users")
	suite.usecase = NewUserUsecase(repository, 10*time.Second)
}

func (suite *userUsecaseSuite) TestCreate_Positive() {
	user := &domain.User{
		Username:  ptr("testuser"),
		Password:  ptr("password"),
	}

	suite.repository.On("Create", mock.Anything, user).Return(nil)

	err := suite.usecase.Create(context.TODO(), user)

	suite.Nil(err)
	suite.repository.AssertExpectations(suite.T())
}

// func (suite *userUsecaseSuite) TestHandleLogin_Positive() {
// 	user := &domain.User{
// 		Username: ptr("testuser"),
// 		Password: ptr("password"),
// 	}

// 	foundUser := &domain.User{
// 		Username:  ptr("testuser"),
// 		Password:  ptr("hashedpassword"),
// 		Email:     ptr("test@example.com"),
// 	}

// 	suite.repository.On("FindByUsername", context.TODO(), *user.Username).Return(foundUser, nil)
// 	infrastructure.On("VerifyPassword", *user.Password, *foundUser.Password).Return(true, "")
// 	infrastructure.On("GenerateJWTToken", foundUser.User_id, *foundUser.Username, *foundUser.Email, foundUser.User_type).Return("token", "refreshToken", nil)

// 	token, refreshToken, err := suite.usecase.HandleLogin(context.TODO(), user)

// 	suite.Nil(err)
// 	suite.Equal("token", token)
// 	suite.Equal("refreshToken", refreshToken)
// 	suite.repository.AssertExpectations(suite.T())
// }

// func (suite *userUsecaseSuite) TestHandleLogin_UserNotFound_Negative() {
// 	user := &domain.User{
// 		Username: ptr("testuser"),
// 		Password: ptr("password"),
// 	}

// 	suite.repository.On("FindByUsername", context.TODO(), *user.Username).Return(nil, errors.New("user not found"))

// 	_, _, err := suite.usecase.HandleLogin(context.TODO(), user)

// 	suite.NotNil(err)
// 	suite.EqualError(err, "user not found")
// 	suite.repository.AssertExpectations(suite.T())
// }

// func (suite *userUsecaseSuite) TestHandleLogin_InvalidPassword_Negative() {
// 	user := &domain.User{
// 		Username: ptr("testuser"),
// 		Password: ptr("password"),
// 	}

// 	foundUser := &domain.User{
// 		Username:  ptr("testuser"),
// 		Password:  ptr("hashedpassword"),
// 		Email:     ptr("test@example.com"),
// 		User_type: "admin",
// 	}

// 	suite.repository.On("FindByUsername", context.TODO(), *user.Username).Return(foundUser, nil)
// 	infrastructure.On("VerifyPassword", *user.Password, *foundUser.Password).Return(false, "invalid password")

// 	_, _, err := suite.usecase.HandleLogin(context.TODO(), user)

// 	suite.NotNil(err)
// 	suite.EqualError(err, "invalid password")
// 	suite.repository.AssertExpectations(suite.T())
// }

// func (suite *userUsecaseSuite) TestUpdate_Positive() {
// 	userID := "user-id"

// 	suite.repository.On("Update", context.TODO(), userID).Return(nil)

// 	err := suite.usecase.Update(context.TODO(), userID)

// 	suite.Nil(err)
// 	suite.repository.AssertExpectations(suite.T())
// }

func TestUserUsecaseSuite(t *testing.T) {
	suite.Run(t, new(userUsecaseSuite))
}

// Helper function to create a pointer to a string
func ptr(s string) *string {
	return &s
}