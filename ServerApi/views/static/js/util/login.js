/* 
* @Author: WhiteWang
* @Date:   2015-04-01 17:43:13
* @Last Modified by:   weifengwang
* @Last Modified time: 2015-04-01 19:15:13
*/

define(function (require, exports, module) {
    var $ = require('jquery');
    var user = require('util/user');
    var hash = require('util/hash');

    var Login = function(){
        var template = $('#login_template').html();
        var $wrapper = $('.wrapper');
        var $loginWrapper = $('.login-wrapper');
        this.show = function(){
            $wrapper.hide();
            var $login = $(template);
            $loginWrapper.html($login);
            var that = this;
            $login.find('.jsLogin').on('click', function(){
                var username = $login.find('input[type=text]').val();
                var password = $login.find('input[type=password]').val();
                var $btn = $(this);
                $btn.parent().addClass('loading');
                user.login(username, password, function(){
                    that.remove();
                    hash.set('work');
                }, function(){
                    $btn.parent().removeClass('loading');
                    $login.find('.error-msg').html('密码错误').show();
                })
            });
            $login.find('input[type=password]').on('focus', function(){
                $login.find('.error-msg').hide();
            })
        }
        this.remove = function(){
            $loginWrapper.html('');
        }
    }

    return new Login();
});