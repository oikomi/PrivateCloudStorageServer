/* 
* @Author: WhiteWang
* @Date:   2015-04-01 15:02:27
* @Last Modified by:   weifengwang
* @Last Modified time: 2015-04-01 16:57:53
*/

define(function (require, exports, module) {
    function getHash(){
        var obj = {
            type: '',
            path: []
        }
        var hash = window.location.hash;
        if(hash.indexOf('#')===0){
            hash = hash.substr(1);
        }
        if(/\/$/.test(hash)){
            hash = hash.substr(0, hash.length-1);
        }
        var t = hash.indexOf('/');
        if(t==-1){
            obj.type = hash;
        } else {
            obj.type = hash.substring(0, t);
            if(hash.substr(t+1)){
                obj.path = hash.substr(t+1).split('/');
            }
        }
        return obj;
    }

    return {
        get: function(){
            return getHash();
        },
        set: function(h){
            window.location.hash = h;
        }
    }
});