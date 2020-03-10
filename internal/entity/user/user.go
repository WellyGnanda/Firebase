package user

import "time"

// JAVA EQUIVALENT -> MODEL

// User object model
type User struct {
	ID           int       `db:"id" json:"user_id"`
	Nip          string    `db:"nip" json:"nip"`
	Nama         string    `db:"nama_lengkap" json:"nama_lengkap"`
	TanggalLahir time.Time `db:"tanggal_lahir" json:"tanggal_lahir"`
	Jabatan      string    `db:"jabatan" json:"jabatan"`
	Email        string    `db:"email" json:"email"`
}

//DataResp Get Data Resp ...
type DataResp struct {
	Data     []User      `json:"data"`
	Metadata interface{} `json:"metadata"`
	Error    interface{} `json:"error"`
}
