package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	pb "github.com/venomuz/project4/UserService/genproto"
	"log"
)

type userRepo struct {
	db *sqlx.DB
}

//NewUserRepo ...
func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db: db}
}

func (r *userRepo) Create(user *pb.Useri) (*pb.Useri, error) {
	UserQuery := `INSERT INTO users(id,first_name,last_name,email,bio,phone_number,type_id,status) VALUES($1,$2,$3,$4,$5,$6,$7,$8)`
	_, err := r.db.Exec(UserQuery, user.Id, user.FirstName, user.LastName, pq.Array(user.Email), user.Bio, pq.Array(user.PhoneNumber), user.TypeId, user.Status)
	if err != nil {
		log.Panicf("%s\n%s", "Error while users to table addresses", err)
	}
	AddressQuery := `INSERT INTO addresses(id,user_id,country,city,district,postal_code) VALUES($1,$2,$3,$4,$5,$6)`

	_, err = r.db.Exec(AddressQuery, user.Address.Id, user.Id, user.Address.Country, user.Address.City, user.Address.District, user.Address.PostalCode)
	if err != nil {
		log.Panicf("%s\n%s", "Error while inserting to table addresses", err)
	}

	return user, nil
}
func (r *userRepo) GetByID(ID string) (*pb.Useri, error) {
	user := pb.Useri{}
	GetUsers := `SELECT id, first_name, last_name, email, bio, phone_number, type_id, status FROM users WHERE id = $1`
	err := r.db.QueryRow(GetUsers, ID).Scan(&user.Id, &user.FirstName, &user.LastName, pq.Array(&user.Email), &user.Bio, pq.Array(&user.PhoneNumber), &user.TypeId, &user.Status)
	if err != nil {
		return nil, err
	}

	addr := pb.Address{}
	GetAddresses := `SELECT id,user_id, city, district, country, postal_code FROM addresses WHERE user_id = $1`
	err = r.db.QueryRow(GetAddresses, user.Id).Scan(&addr.Id, &addr.UserId, &addr.City, &addr.District, &addr.Country, &addr.PostalCode)
	if err != nil {
		return nil, err
	}
	user.Address = &addr
	return &user, nil
}
func (r *userRepo) DeleteByID(ID string) (*pb.GetIdFromUserID, error) {
	_, err := r.db.Exec(`DELETE  FROM users WHERE id = $1`, ID)
	if err != nil {
		log.Panicf("%s\n%s", "Error while deleteing data from table users", err)
	}
	id := pb.GetIdFromUserID{}

	return &id, nil
}
func (r *userRepo) GetAllUserFromDb(empty *pb.Empty) (*pb.AllUser, error) {
	var userss pb.AllUser
	user := pb.Useri{}
	GetUsers := `SELECT id, first_name, last_name, email, bio, phone_number, type_id, status FROM users;`
	rows, err := r.db.Query(GetUsers)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, pq.Array(&user.Email), &user.Bio, pq.Array(&user.PhoneNumber), &user.TypeId, &user.Status)
		if err != nil {
			return nil, err
		}
		addr := pb.Address{}
		GetAddresses := `SELECT id,user_id, city, district, country, postal_code FROM addresses WHERE user_id = $1`
		err = r.db.QueryRow(GetAddresses, user.Id).Scan(&addr.Id, &addr.UserId, &addr.City, &addr.District, &addr.Country, &addr.PostalCode)
		if err != nil {
			return nil, err
		}
		user.Address = &addr
	}
	userss.Users = append(userss.Users, &user)

	return &userss, nil
}
