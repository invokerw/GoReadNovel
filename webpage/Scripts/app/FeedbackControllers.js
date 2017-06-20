"use strict";

/* controllers module */

angular.module("BsTableDirective.FeedbackControllers", ["BsTableDirective.Services","ui.bootstrap"])
    .controller("FeedbackCtrl", ["$scope", "BootswatchService","$http", 
        function ($scope, BootswatchService, $http) {
        // scope models
        $scope.page = { Css: "/statics/bootstrap/css/bootstrap.min.css" };
        $scope.contactList = [];
        $scope.progress = { Ready: false };
        // get themes from bootswatch
        BootswatchService.GetAll().success(function (result) {
            $scope.page.Themes = result.themes;
        });
        // get data for bs-table
        GenerateData();
        // show function for bs-table
        $scope.Show = function (contact) {
            alert(JSON.stringify(contact));
        };
        // Solved function for bs-table
        $scope.Solved = function (contact) {
            //alert(JSON.stringify(contact));
            $http({
                method: 'GET',
                url: '/SolvedUserFeedbackJson',
                params:{
                    contactid:contact.feedbackid
                }
            }).then(function successCallback(response) {
                // 请求成功执行代码
                console.log('response.data = ',response.data); 
                if (response.data.code == 1) {
                    GenerateData()
                }   
                // hide progress
            }, function errorCallback(response) {
            // 请求失败执行代码
                alert("Http Get Error:SolvedUserFeedbackJson");

            });
        };
        // remove function for bs-table
        $scope.Remove = function (contact) {
            //alert(JSON.stringify(contact));
            $http({
                method: 'GET',
                url: '/DelAUserFeedbackJson',
                params:{
                    contactid:contact.feedbackid
                }
            }).then(function successCallback(response) {
                // 请求成功执行代码
                console.log('response.data = ',response.data); 
                if (response.data.code == 1) {
                    GenerateData()
                }  
                // hide progress
            }, function errorCallback(response) {
            // 请求失败执行代码
                alert("Http Get Error:DelAUserFeedbackJson");

            });
        };
        // generate data
        function GenerateData() {
            $http({
                method: 'GET',
                url: '/GetUserFeedbackJson',
                params:{
                    feedbacktype:-1,
                    solved:-1
                }
            }).then(function successCallback(response) {
                // 请求成功执行代码
                console.log('response.data = ',response.data);   
                // hide progress
                $scope.progress = { Ready: true };
                $scope.contactList = response.data.ret;
                //return dataList;
            }, function errorCallback(response) {
            // 请求失败执行代码
                alert("Http Get Error:GetUserFeedbackJson");

            });
        }
        // init socialite jquery plugin
        Socialite.load();
    }]);
    