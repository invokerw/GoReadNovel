"use strict";

/* app module */

angular.module("BsTableDirective", ["ngRoute", "bsTable","BsTableDirective.FeedbackControllers",
    "BsTableDirective.NovelControllers", "BsTableDirective.UserControllers",
    "BsTableDirective.SpiderControllers","BsTableDirective.SysNoticeControllers",
    "BsTableDirective.Services"])
    .config(["$routeProvider", function ($routeProvider) {
        $routeProvider.when("/", {
            templateUrl: "/Content/templates/feedback.html",
            controller: "FeedbackCtrl"
        })
        .when("/ShowNovelInfo", {
            templateUrl: "/Content/templates/noveltable.html",
            controller: "NovelInfoCtrl"
        })
        .when("/ShowUserInfo", {
            templateUrl: "/Content/templates/usertable.html",
            controller: "UserCtrl"
        })
        .when("/SpiderInfo", {
            templateUrl: "/Content/templates/spiderinfo.html",
            controller: "SpiderCtrl"
        })
        .when("/SysNoticeInfo", {
            templateUrl: "/Content/templates/sysnotice.html",
            controller: "SysNoticeCtrl"
        })
        .otherwise({ redirectTo: "/" });
    }])
    .filter('removal',function(){
    	return function(text) {
    	//var str = "http://www.huanyue123.com/book"; 
    	//str.length = 30;
        return text.substring(30);
    	}
    })
    .filter('solve',function(){
        return function(text) {
        var str = "";
        if (text == 0) {
            str = "未解决";
        }else if(text == 1){
            str = "已解决";
        }else{
            str = "不应该出现这句话";
        }
        return str;
        }
    })
    .filter('feedbacktype',function(){
        return function(text) {
        var str = "";
        if (text == 0) {
            str = "书籍问题";
        }else if(text == 1){
            str = "操作问题";
        }else if(text == 2){
            str = "其他问题";
        }else{
            str = "未知问题";
        }
        return str;
        }
    })
    .filter('timechange',function(){
        return function(text) {
            console.log(text)
            return new Date(parseInt(text) * 1000).toLocaleString().substr(0,17)
        }
    })
    .filter('gender',function(){
        return function(text) {
        var str = "";
        if (text == 2) {
            str = "女";
        }else if(text == 1){
            str = "男";
        }else{
            str = "未知性别";
        }
        return str;
        }
    });
                
