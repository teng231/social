const CLOUDINARY_NAME = 'telosma'
const CLOUDINARY_API_KEY = '185362918485478'
const CLOUDINARY_API_SECRET = '97CGQGl4iIQ35AOWYpu_3u2S564'

const cloudinary = require('cloudinary')
const Promise = require('bluebird')
var fs = require('fs')
cloudinary.config({
  cloud_name: CLOUDINARY_NAME,
  api_key: CLOUDINARY_API_KEY,
  api_secret: CLOUDINARY_API_SECRET
})

function uploadWithHttpUrl(url, folder = 'commons', tags) {
  return new Promise(function (resolve, reject) {
    cloudinary.v2.uploader.upload(url,
      {
        folder: `${folder}`,
        tags
      },
      function (error, result) {
        if (error) {
          reject(error)
          return
        }
        let {
          public_id,
          width,
          height,
          tags,
          bytes,
          format,
          url,
          original_filename
        } = result
        resolve({
          public_id,
          width,
          height,
          tags,
          bytes,
          format,
          url,
          original_filename
        })
      })
  })
}

// uploadWithHttpUrl('https://images.unsplash.com/photo-1518889778-5111daad1bda?ixlib=rb-0.3.5&q=85&fm=jpg&crop=entropy&cs=srgb&ixid=eyJhcHBfaWQiOjEyMDd9&s=6ecf48c54c0916adaddba10b0336617f', 'test', [])

function uploadWithFile(path, tags) {
  return new Promise(function (resolve, reject) {
    var stream = cloudinary.v2.uploader.upload_stream({ tags }, function (error, result) {
      // console.log(result)
      result = {
        public_id: result.public_id,
        width: result.width,
        height: result.height,
        bytes: result.bytes,
        url: result.url,
        format: result.format
      }
      resolve(result)
      return
    })
    var file_reader = fs.createReadStream(path).pipe(stream)
  })
}

function readFilesName(dirname, cb) {
  fs.readdir(dirname, function (err, filenames) {
    if (err) {
      onError(err);
      return;
    }
    cb(filenames)
    // console.log(JSON.stringify(filenames))
  })
}
let axios = require('axios')
let faker = require('faker')
let chunk = require('lodash/chunk')
let folder = 'nganquynh'
let uid = '5ad1e1bad552310a06900ea5'
let content = "Tuổi trẻ ấy, thật vất vả và buồn cười."
let url = `http://local.tenm.cf:4444/post/${uid}/create`


readFilesName(`../albums/${folder}`, (filenames) => {
  filenames = filenames.filter(v => v !== '.DS_Store')
  let uploadReq = filenames.map(v => {
    return uploadWithFile(`../albums/${folder}/${v}`, [
      "tree",
      "woodland",
      "forest",
      "female",
      "long hair",
      "woman",
      "girl",
      "sport",
      "bag",
      "path",
      "track",
      "nature",
      "green",
      "person",
      "wilderness"
    ])
  })
  Promise.all(uploadReq)
    .then(data => {
      data = chunk(data, 3)
      return data.map(v => {
        return createPost(v)
      })
    })
    .then(arrReq => {
      Promise.all(arrReq).then(data => {
        console.log('done')
      })
    })
})

function createPost(media) {
  let tags = JSON.stringify([
    "tree",
    "woodland",
    "forest",
    "female",
    "long hair",
    "woman",
    "girl",
    "sport",
    "bag",
    "path",
    "track",
    "nature",
    "green",
    "person",
    "wilderness"
  ])
  media = JSON.stringify(media)
  return axios.default.post(url,
    `user_id=${uid}&medias=${media}&tags=${tags}&content=${content}`, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8'
      }
    })
}

module.exports = uploadWithHttpUrl