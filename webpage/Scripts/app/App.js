"use strict";

/* app module */

angular.module("BsTableDirective", ["ngRoute", "bsTable", "BsTableDirective.NovelControllers", 
	"BsTableDirective.UserControllers","BsTableDirective.SpiderControllers","BsTableDirective.Services"])
    .config(["$routeProvider", function ($routeProvider) {
        $routeProvider.when("/", {
            templateUrl: "/Content/templates/example.html",
            controller: "ExampleCtrl"
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
        .otherwise({ redirectTo: "/" });
    }])
    .filter('removal',function(){
    	return function(text) {
    	//var str = "http://www.huanyue123.com/book"; 
    	//str.length = 30;
        return text.substring(30);
    	}
    });
