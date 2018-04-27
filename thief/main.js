var request = require('request')
let db = require('./src/writer')
let Promise = require('bluebird')
const queryString = require('query-string');
let qs = (o) => {
  return queryString.stringify(o)
}

let params = {
  query: 'flower',
  per_page: 30
}

var options = (page, token) => ({
  url: `https://unsplash.com/napi/search/photos?${qs(params)}&page=${page}`,
  headers: {
    'authorization': `Bearer ${token}`
  }
})

function convertObject(list) {
  return list.map(l => {
    return {
      public_id: `natural${l.id}`,
      url: l.urls.full,
      width: l.width,
      height: l.height,
      description: l.description,
      tags: l.photo_tags.map(v => { return v.title })
    }
  })
}

function callback(error, response, body) {
  if (!error && response.statusCode == 200) {
    let data = convertObject(JSON.parse(body).results)
    db.create(data).then(res => {
      console.log(res.length)
      return
    })
  }
}

let runSave = (initPage = 1, token) => {
  let maxRun = 99 + Number(initPage)
  initPage = Number(initPage)

  for (let i = initPage; i < maxRun; i++) {
    request(options(i, token), callback)
  }
}

let requestToken = process.argv[2]
let initPage = process.argv[3]
if (!initPage || !requestToken) return new Error('missing params')
console.log('run', initPage, requestToken)

runSave(initPage, requestToken)