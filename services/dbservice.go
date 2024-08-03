package dbservice

import (
	"context"
	"fmt"
	"http_request/class-management/config"
	"http_request/class-management/models"
	"log"

	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

type DatabaseService struct {
	Database driver.Database
}

func InitDatabase() DatabaseService {
	print("Initializing database\n")
	appConfig := config.LoadAppConfig()

	// Create an HTTP connection to the database
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{"http://localhost:" + appConfig.DbPort},
	})
	if err != nil {
		fmt.Println("Failed to create HTTP connection to database")
		log.Fatal(err)
	}
	// Create a client
	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(appConfig.DbUser, appConfig.DbPassword),
	})

	if err != nil {
		fmt.Println("Failed to create client")
		log.Fatal(err)

	}

	ctx := context.Background()

	db, err := client.Database(ctx, appConfig.DbName)

	if err != nil {
		// Create the database if it does not exist
		db, err = client.CreateDatabase(ctx, appConfig.DbName, nil)
		if err != nil {
			fmt.Println("Failed to create database")
			log.Fatal(err)
		}

		fmt.Println("Database created")
	} else {
		fmt.Println("Database already exists")
	}

	collections := []string{"student", "teacher", "user", "class"}

	// Check if collections exist in the database, if not, create them
	for _, collection := range collections {
		exists, err := db.CollectionExists(ctx, collection)
		if err != nil {
			log.Fatal(err)

		}
		if !exists {
			_, err := db.CreateCollection(ctx, collection, nil)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Collection %s created\n", collection)
		} else {
			fmt.Printf("Collection %s already exists\n", collection)
		}
	}

	fmt.Println("Database initialized successfully")
	return DatabaseService{Database: db}
}

func (dbs *DatabaseService) GetUserByUsername(username string) (*models.User, error) {
	ctx := context.Background()

	collection, err := dbs.Database.Collection(ctx, "user")
	if err != nil {
		return nil, err
	}

	// get user from database by username
	query := fmt.Sprintf("FOR u IN user FILTER u._key == '%s' RETURN u", username)
	bindVars := map[string]interface{}{}

	cursor, err := collection.Database().Query(ctx, query, bindVars)

	if err != nil {
		if driver.IsNotFoundGeneral(err) {
			return nil, nil
		}

		return nil, err
	}

	defer cursor.Close()

	var user models.User
	_, err = cursor.ReadDocument(ctx, &user)

	if err != nil {
		if driver.IsNoMoreDocuments(err) {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}
func (dbs *DatabaseService) CreateUser(user *models.User) error {
	ctx := context.Background()
	collection, err := dbs.Database.Collection(ctx, "user")
	if err != nil {
		return err
	}
	_, err = collection.CreateDocument(context.Background(), user)
	return err
}

func (dbs *DatabaseService) CreateStudent(student *models.Student) (*models.Student, error) {
	ctx := context.Background()
	collection, err := dbs.Database.Collection(ctx, "student")
	if err != nil {
		return nil, err
	}
	meta, err := collection.CreateDocument(ctx, student)

	student.Key = meta.Key
	return student, err
}

func (dbs *DatabaseService) GetStudentByID(id string) (*models.Student, error) {
	ctx := context.Background()
	collection, err := dbs.Database.Collection(ctx, "student")
	if err != nil {
		return nil, err
	}

	var student models.Student
	_, err = collection.ReadDocument(ctx, id, &student)
	if err != nil {
		if driver.IsNotFoundGeneral(err) {
			return nil, nil
		}
		return nil, err
	}

	return &student, nil
}

func (dbs *DatabaseService) UpdateStudent(id string, student *models.Student) error {
	ctx := context.Background()
	collection, err := dbs.Database.Collection(ctx, "student")
	if err != nil {
		return err
	}

	_, err = collection.UpdateDocument(ctx, id, student)
	if err != nil {
		return err
	}

	return nil
}

func (dbs *DatabaseService) DeleteStudent(id string) error {
	ctx := context.Background()
	collection, err := dbs.Database.Collection(ctx, "student")
	if err != nil {
		return err
	}

	_, err = collection.RemoveDocument(ctx, id)
	return err
}

func (dbs *DatabaseService) CreateTeacher(teacher *models.Teacher) (*models.Teacher, error) {
	ctx := context.Background()
	collection, err := dbs.Database.Collection(ctx, "teacher")
	if err != nil {
		return nil, err
	}
	meta, err := collection.CreateDocument(ctx, teacher)

	teacher.ID = meta.Key
	return teacher, err
}

func (dbs *DatabaseService) GetTeacherByID(id string) (*models.Teacher, error) {
	ctx := context.Background()
	collection, err := dbs.Database.Collection(ctx, "teacher")
	if err != nil {
		return nil, err
	}

	var teacher models.Teacher
	meta, err := collection.ReadDocument(ctx, id, &teacher)
	if err != nil {
		if driver.IsNotFoundGeneral(err) {
			return nil, nil
		}
		return nil, err
	}

	teacher.ID = meta.Key
	return &teacher, nil
}

func (dbs *DatabaseService) UpdateTeacher(id string, teacher *models.Teacher) error {
	ctx := context.Background()
	collection, err := dbs.Database.Collection(ctx, "teacher")
	if err != nil {
		return err
	}

	meta, err := collection.UpdateDocument(ctx, id, teacher)
	if err != nil {
		return err
	}

	teacher.ID = meta.Key
	return nil
}

func (dbs *DatabaseService) DeleteTeacher(id string) error {
	ctx := context.Background()
	collection, err := dbs.Database.Collection(ctx, "teacher")
	if err != nil {
		return err
	}

	_, err = collection.RemoveDocument(ctx, id)
	return err
}

func (dbs *DatabaseService) CreateClass(class *models.Class) error {
	ctx := context.Background()
	collection, err := dbs.Database.Collection(ctx, "class")
	if err != nil {
		return err
	}
	_, err = collection.CreateDocument(ctx, class)
	return err
}

func (dbs *DatabaseService) UpdateClassTeacher(classID string, teacherID string) error {
	ctx := context.Background()
	collection, err := dbs.Database.Collection(ctx, "class")
	if err != nil {
		return err
	}

	_, err = collection.UpdateDocument(ctx, classID, map[string]interface{}{
		"teacher_id": teacherID,
	})
	return err
}

func (dbs *DatabaseService) AddStudentToClass(classID string, studentID string) error {
	ctx := context.Background()
	collection, err := dbs.Database.Collection(ctx, "class")
	if err != nil {
		return err
	}

	var class models.Class
	_, err = collection.ReadDocument(ctx, classID, &class)
	if err != nil {
		return err
	}

	class.StudentIDs = append(class.StudentIDs, studentID)

	_, err = collection.UpdateDocument(ctx, classID, map[string]interface{}{
		"student_ids": class.StudentIDs,
	})
	return err
}

func (dbs *DatabaseService) RemoveStudentFromClass(classID string, studentID string) error {
	ctx := context.Background()
	collection, err := dbs.Database.Collection(ctx, "class")
	if err != nil {
		return err
	}

	var class models.Class
	_, err = collection.ReadDocument(ctx, classID, &class)
	if err != nil {
		return err
	}

	// Remove studentID from StudentIDs
	for i, id := range class.StudentIDs {
		if id == studentID {
			class.StudentIDs = append(class.StudentIDs[:i], class.StudentIDs[i+1:]...)
			break
		}
	}

	_, err = collection.UpdateDocument(ctx, classID, map[string]interface{}{
		"student_ids": class.StudentIDs,
	})
	return err
}

func (dbs *DatabaseService) GetClassByID(classID string) (*models.Class, error) {
	ctx := context.Background()
	collection, err := dbs.Database.Collection(ctx, "class")
	if err != nil {
		return nil, err
	}

	var class models.Class
	_, err = collection.ReadDocument(ctx, classID, &class)
	if err != nil {
		if driver.IsNotFoundGeneral(err) {
			return nil, nil
		}
		return nil, err
	}

	// Fetch Teacher
	if class.TeacherID != "" {
		teacherCollection, err := dbs.Database.Collection(ctx, "teacher")
		if err != nil {
			return nil, err
		}
		var teacher models.Teacher
		_, err = teacherCollection.ReadDocument(ctx, class.TeacherID, &teacher)
		if err != nil {
			if driver.IsNotFoundGeneral(err) {
				class.Teacher = nil
			} else {
				return nil, err
			}
		} else {
			class.Teacher = &teacher
		}
	}

	// Fetch Students
	if len(class.StudentIDs) > 0 {
		studentCollection, err := dbs.Database.Collection(ctx, "student")
		if err != nil {
			return nil, err
		}
		var students []*models.Student
		for _, studentID := range class.StudentIDs {
			var student models.Student
			_, err = studentCollection.ReadDocument(ctx, studentID, &student)
			if err != nil {
				if driver.IsNotFoundGeneral(err) {
					continue
				}
				return nil, err
			}
			students = append(students, &student)
		}
		class.Students = students
	}

	return &class, nil
}

func (dbs *DatabaseService) DeleteClassByID(classID string) error {
	ctx := context.Background()
	collection, err := dbs.Database.Collection(ctx, "class")
	if err != nil {
		return err
	}

	_, err = collection.RemoveDocument(ctx, classID)
	return err
}
