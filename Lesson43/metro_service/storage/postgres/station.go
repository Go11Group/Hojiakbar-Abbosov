package postgres

import (
	"database/sql"

	"github.com/Go11Group/Hojiakbar-Abbosov/Lesson43/metro_service/models"
	"github.com/google/uuid"
)

type StationRepo struct {
	Db *sql.DB
}

func NewStationRepo(db *sql.DB) *StationRepo {
	return &StationRepo{Db: db}
}

func (s *StationRepo) Create(station *models.CreateStation) error {
	_, err := s.Db.Exec("insert into station(id, name) values ($1, $2)",
		uuid.NewString(), station.Name)
	return err
}

func (s *StationRepo) GetById(id string) (*models.Station, error) {
	var station = models.Station{Id: id}

	err := s.Db.QueryRow("select name from station where id = $1", id).
		Scan(&station.Name)
	if err != nil {
		return nil, err
	}

	return &station, nil
}

func (s *StationRepo) Update(station *models.UpdateStation) error {
	_, err := s.Db.Exec("update station set name = $1 where id = $2",
		station.Name, station.Id)
	return err
}

func (s *StationRepo) Delete(id string) error {
	_, err := s.Db.Exec("delete from station where id = $1", id)
	return err
}
