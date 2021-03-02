var express = require('express');
var router = express.Router();
var co = require('co')
var queryChaincode = require('../controllers/queryChaincode');


/* GET home page. */
router.get('/', function(req, res, next) {
  res.render('index', { title: 'Express' });
});

/**
 * 合约查询，查询余额
 */
router.get('/query', function(req, res, next){
  let params = req.query;
  let args = [params.id];
  let fcn = 'query';
  co(function *(){
    let result = yield queryChaincode.query(fcn, args);
    if(result.code == 200 || result.code == null){
      res.code = '200';
      res.send(result.msg);
  }else{
      res.code = '500';
      res.send(res.msg);
  }
  })
})

module.exports = router;
