let request = require('request')


const getData = async (endPoint) => {
  await request
    .get(endPoint)
    .on('response', function (response) {
      console.log(response.statusCode) // 200
      console.log(response.headers['content-type']) // 'image/png'
    })
}