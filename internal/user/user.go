package user

type User struct {
	ID    uint `gorm:"primaryKey;autoIncrement"`
	Email string
	Age   int
}
