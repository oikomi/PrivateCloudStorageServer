/* 
* @Author: WhiteWang
* @Date:   2015-04-01 17:19:16
* @Last Modified by:   weifengwang
* @Last Modified time: 2015-04-01 17:59:15
*/

define(function (require, exports, module){
    var cookie = require('util/cookie');
    var ajax = require('util/ajax');
    var api = require('api');
    var hash = require('util/hash');

    var User = function(){
        this.username = cookie.read('username');
        this.token = cookie.read('token');
        this.isLogin = function(){
            var username = this.username || cookie.read('username');
            var token = this.token || cookie.read('token');
            if(username && token){
                return true;
            } else {
                return false;
            }
        }
        this.login = function(username, password, callback, fail){
            var that = this;
            ajax({
                url: api.login,
                type: 'POST',
                data: {user_name: username, password:password},
                success: function(data){
                    if(data.status==0 && data.token){
                        cookie.write('username', username, {duration:7});
                        cookie.write('token', data.token, {duration:7});
                        that.username = username;
                        that.token = data.token;
                        callback && callback();
                    } else {
                        fail && fail();
                    }
                },
                error: function(){
                    fail && fail();
                }
            })
        }
        this.logout = function(){
            cookie.remove('username');
            cookie.remove('token');
            this.username = null;
            this.token = null;
            hash.set('login');
        }
    };
    return new User();
});