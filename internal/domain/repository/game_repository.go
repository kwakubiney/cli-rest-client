package repository
import (

	"github.com/kwakubiney/canonical-take-home/internal/domain/model"
	"gorm.io/gorm"
	"fmt"
	"log"
)
type GameNotFoundError struct {
    msg string
}

func (e *GameNotFoundError) Error() string { return e.msg }

type GameRepository struct {
	db *gorm.DB
}

func NewGameRepository(db *gorm.DB) *GameRepository {
	return &GameRepository{
		db,
	}
}

func (g *GameRepository) CreateGame(game model.Game) error {
	return g.db.Create(&game).Error
}

func (g *GameRepository) UpdateGameByTitle(title string, game model.Game) error {
	db :=  g.db.Model(model.Game{}).Where("title = ?", title).Find(&game)
	if db.RowsAffected == 0 {
		return  &GameNotFoundError{"no record found for title given"}
	}
	err := db.Updates(&game).Error
	if err != nil {
		log.Println(db.Error)
		return err
	}
	return nil
}

func (g *GameRepository) FilterGame(by string, where string) (*model.Game, error) {
	var game model.Game
	db := g.db.Where(fmt.Sprintf("%s = ?", by), where).Find(&game)
	if db.RowsAffected == 0 {
		return &game, &GameNotFoundError{"no record found for title given"}
	}
	if db.Error != nil {
		log.Println(db.Error)
		return nil, db.Error
	}
	return &game, db.Error
}

