var mongoose = require('mongoose')

var db = mongoose.connection;

db.on('error', console.error)

mongoose.connect('mongodb://user:123@ds151222.mlab.com:51222/dbbase')

var Medias = new mongoose.Schema({
  public_id: String,
  url: String,
  width: Number,
  width: Number,
  description: String,
  tags: [String]
})
// tao model smarjob tương ứng với schema đã khai báo bên trên
module.exports = mongoose.model('Medias', Medias, 'flower')
