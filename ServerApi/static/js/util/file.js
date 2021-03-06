/* 
* @Author: White
* @Email: weifengwang@pptv.com
* @Date:   2015-04-01 19:48:55
* @Last Modified time: 2015-04-12 17:11:51
*/

define(function(require, exports, module) {
    var $ = require('jquery');
    var _ = require('underscore');
    var adaptor = require('util/adaptor');
    var hash = require('util/hash');
    var api = require('api');

    var template = _.template($('#file_template').html());
    var $space = $('.file-con');
    function getIcon(obj){
        var icon = 'unknow-icon';
        if(obj.type == 'dir'){
            icon =  'folder-icon';
        } else {
            var type = /\.[^\.]+$/.exec(obj.name);
            if(type){
                type = type[0];
            }
            switch(type){
                case '.pdf':
                    icon = 'pdf-icon';break;
                case '.doc':
                    icon = 'word-icon';break;
                case '.xls':
                    icon = 'xls-icon';break;
                case '.zip':
                    icon = 'zip-icon';break;
                default:
                    icon = 'unknow-icon';break;
            }
        }
        return icon;
    }
    function getListDom(obj){
        var $li = $(template({
            fileIcon: getIcon(obj),
            fileName: obj.name
        }));
        $li.on('mouseenter', function(){
            $(this).addClass('hover');
        }).on('mouseleave', function(){
            $(this).removeClass('hover');
        }).on('click', function(){
            $(this).siblings().removeClass('click');
            $(this).addClass('click');
            adaptor.set(obj, $(this));
        }).on('dblclick', function(){
            if(obj.type == 'dir'){
                hash.set('work'+obj.path);
            } else {
                window.open(api.download+obj.path);
            }
        });
        return $li;
    }
    function rePath(obj){
        while(obj.path.substr(0,1)=='/' && obj.path.substr(1,1)=='/'){
            obj.path = obj.path.substr(1);
        }
        return obj;
    }
    function File(){
        this.init = function(data){
            $space.children('ul').html('');
            data.sort(function(a, b){
                if(a.name && b.name && a.name<b.name){
                    return -1;
                } else {
                    return 1;
                }
            });
            var dArr = [];
            var fArr = [];
            for(var i=0; i<data.length; i++){
                var obj = data[i];
                if(!!obj.status){
                    continue;
                }
                if(obj.type == 'dir'){
                    dArr.push(obj);
                } else {
                    fArr.push(obj);
                }
            }
            data = dArr.concat(fArr);
            for(var i=0; i<data.length; i++){
                var obj = data[i];
                if(!!obj.status){
                    continue;
                }
                this.add(obj);
            }
        }
        this.add = function(obj){
            $space.children('ul').append(getListDom(rePath(obj)));
        }
        this.prepend = function(obj){
            $space.children('ul').prepend(getListDom(rePath(obj)));
        }
    }

    return new File();
});