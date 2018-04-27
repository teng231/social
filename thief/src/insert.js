const CLOUDINARY_NAME = 'telosma'
const CLOUDINARY_API_KEY = '185362918485478'
const CLOUDINARY_API_SECRET = '97CGQGl4iIQ35AOWYpu_3u2S564'

const cloudinary = require('cloudinary')
const Promise = require('bluebird')

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
          public_id, width, height, tags,
          bytes, format, url, original_filename
        } = result
        resolve({
          public_id, width, height, tags,
          bytes, format, url, original_filename
        })
      })
  })
}


module.exports = uploadWithHttpUrl