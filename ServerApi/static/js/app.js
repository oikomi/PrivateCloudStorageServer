/* 
 * @Author: White
 * @Email: weifengwang@pptv.com
 * @Date:   2015-03-28 13:58:02
 * @Last Modified time: 2015-04-13 18:52:52
 */
define(function (require, exports, module) {
    var $ = require('jquery');
    require('lib/jquery.hashchange')($);
    var hash = require('util/hash');
    var user = require('util/user');
    var login = require('util/login');
    var work = require('util/work');
    var upload = require('util/upload');
    require('util/tool');

    $(window).hashchange(function(){
        var h = hash.get();
        switch (h.type){
            case 'login':
                if(user.isLogin()){
                    hash.set('work');
                } else {
                    login.show();
                }
                break;
            case 'work':
                if(user.isLogin()){
                    $('.wrapper').show();
                    $('.userArea span').html(user.username);
                    work.init(h.path);
                } else {
                    hash.set('login');
                }
                break;
            default:
                hash.set('work');
                break;
        }
    });
    $('.jsLogout').on('click', function(){
        user.logout();
    })
    $('.jsBack').on('click', function(){
        if($(this).hasClass('disable')){
            return;
        }
        var path = hash.get().path;
        path.pop();
        hash.set('work/'+path.join('/'));
    })
    $('.upload a').on('click', function(ev){
        upload.init();
    })
    $('body').on('click', '.popup .close, .popup .jsCancel', function(){
        $(this).parents('.popup').remove();
    })
    $('.file-con').height(getfileheight());
    $(window).on('resize', function(){
        $('.file-con').height(getfileheight());
    })
    $(window).hashchange();
    function getfileheight(){
        var windowH = $(window).height();
        return windowH-50-45;
    }
});