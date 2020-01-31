//SPDX-License-Identifier: Apache-2.0

var user = require('./controller.js');

module.exports = function(app){

  app.get('/get_user/:id', function(req, res){
    tuna.get_user(req, res);
  });
  app.get('/create_user/:tuna', function(req, res){
    tuna.create_user(req, res);
  });
  app.get('/list_user', function(req, res){
    tuna.list_tuna(req, res);
  });
}
