/* 
* @Author: White
* @Email: weifengwang@pptv.com
* @Date:   2015-04-07 23:37:35
* @Last Modified time: 2015-04-08 00:35:06
*/

define(function(require, exports, module) {
    var $ = require('jquery');
    var api = require('api');
    var user = require('util/user');
    var hash = require('util/hash');
    var file = require('util/file');

    var tp_upload = '<div class="popup">'+
                        '<div class="pop-upload">'+
                            '<form method="post" enctype="multipart/form-data" target="uploadtg">'+
                                '<input type="file" name="file" />'+
                                '<div class="upsubmit">'+
                                    '<input type="submit" value="上传" class="load-hidden" />'+
                                    '<img class="loading-img" src="images/loading.gif" />'+
                                '</div>'+
                            '</form>'+
                            '<iframe name="uploadtg" id="uploadtg" width="0" height="0"></iframe>'+
                            '<a href="javascript:;" class="close">x</a>'+
                        '</div>'+
                    '</div>';
    function Upload(){
        this.init = function(){
            var $pop = $(tp_upload);
            var fileobj = {
                status:0,
                type:'file'
            };
            $('body').append($pop);
            $pop.find('input[type=file]').on('change',function(){
                var $this = $(this);
                var filename = getFileName($this.val());
                var path = '/';
                var urlpath = hash.get().path;
                if(urlpath && urlpath.length!=0){
                    path = path+urlpath.join('/')+'/'+filename;
                } else {
                    path = path+filename;
                }
                fileobj.name = filename;
                fileobj.path = path;
                var actionurl = api.upload+'&token='+user.token+'&path='+path;
                $pop.find('form').attr('action', actionurl);
            });
            $pop.find('input[type=submit]').on('click',function(){
                $(this).parent().addClass('loading')
            })
            $pop.find('#uploadtg').on('load', function(){
                $pop.remove();
                file.add(fileobj);
            })
        }
    }
    function getFileName(str){
        var reg = /[^\\\/]*[\\\/]+/g;
        str = str.replace(reg,'');
        return str;
    }

    return new Upload();
});