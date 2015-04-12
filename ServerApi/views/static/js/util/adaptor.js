/* 
* @Author: White
* @Email: weifengwang@pptv.com
* @Date:   2015-04-01 21:44:59
* @Last Modified time: 2015-04-06 23:34:25
*/

define(function(require, exports, module) {
    var $ = require('jquery');
    function Adaptor(){
        var $jsCopy = $('.jsCopy'),
            $jsMove = $('.jsMove'),
            $jsDelete = $('.jsDelete'),
            $jsRename = $('.jsRename'),
            $jsDownload = $('.jsDownload');
        var fileObj = null;
        this.set = function(obj, el){
            $jsCopy.removeClass('disable');
            $jsMove.removeClass('disable');
            $jsDelete.removeClass('disable');
            $jsRename.removeClass('disable');
            $jsDownload.removeClass('disable');
            if(obj && typeof obj=='object'){
                if(obj.type == 'dir'){
                    $jsDownload.addClass('disable');
                    $jsCopy.addClass('disable');
                } else {
                }
                fileObj = {};
                fileObj.$el = el;
                fileObj.obj = obj;
            } else {
                $jsCopy.addClass('disable');
                $jsMove.addClass('disable');
                $jsDelete.addClass('disable');
                $jsRename.addClass('disable');
                $jsDownload.addClass('disable');
                fileObj = null;
            }
        }
        this.destroy = function(){
            this.set(null);
        }
        this.get = function(){
            return fileObj;
        }
    }

    return new Adaptor();
});