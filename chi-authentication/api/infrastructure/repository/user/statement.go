package user

var insertUser = `INSERT INTO users(id, name, mail, imagePath, pass) VALUE(?,?,?,?,?)`
var selectUserByMail = `SELECT id, name, mail, imagePath FROM users WHERE mail=?`
var selectUserPassByMail = `SELECT pass FROM users WHERE mail=?`
