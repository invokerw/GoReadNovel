"use strict";

/* controllers module */

angular.module("BsTableDirective.Controllers", ["BsTableDirective.Services"])
    .controller("ExampleCtrl", ["$scope", "BootswatchService", function ($scope, BootswatchService) {
        // scope models
        $scope.page = { Css: "/statics/bootstrap/css/bootstrap.min.css" };
        $scope.contactList = [];
        $scope.progress = { Ready: false };
        // get themes from bootswatch
        BootswatchService.GetAll().success(function (result) {
            // set themes to scope
            $scope.page.Themes = result.themes;
            // hide progress
            $scope.progress = { Ready: true };
        });
        // get data for bs-table
        $scope.contactList = GenerateData(20);
        // show function for bs-table
        $scope.Show = function (contact) {
            alert(JSON.stringify(contact));
        };
        // edit function for bs-table
        $scope.Edit = function (contact) {
            alert(JSON.stringify(contact));
        };
        // remove function for bs-table
        $scope.Remove = function (contact) {
            alert(JSON.stringify(contact));
        };
        // generate data
        function GenerateData(count) {
            var dataList = [];
            for (var i = 0; i < count; i++)
            {
                // create object
                var item = {
                    FirstName: "FirstName " + i,
                    LastName: "LastName " + i,
                    BornDate: "2014-07-10"
                };
                // add to list
                dataList.push(item);
            }
            return dataList;
        }
        // init socialite jquery plugin
        Socialite.load();
    }])
    .controller("NovelInfoCtrl", ["$scope", "BootswatchService", "$http", function ($scope, BootswatchService, $http) {
        // scope models
        $scope.page = { Css: "/statics/bootstrap/css/bootstrap.min.css" };
        $scope.contactList = [];
        $scope.progress = { Ready: false };
        // get themes from bootswatch
        BootswatchService.GetAll().success(function (result) {
            // set themes to scope
            $scope.page.Themes = result.themes;
            // hide progress
            //$scope.progress = { Ready: true };
        });
        // get data for bs-table
        GenerateData(0,100);

        // show function for bs-table
        $scope.Show = function (contact) {
            alert(JSON.stringify(contact));
        };
        // edit function for bs-table
        $scope.Edit = function (contact) {
            alert(JSON.stringify(contact));
        };
        // remove function for bs-table
        $scope.Remove = function (contact) {
            alert(JSON.stringify(contact));
        };
        // generate data
        function GenerateData(begin,count) {
            var dataList = [];
            $http({
                method: 'GET',
                url: '/GetTNovelTableInfoJson',
                params: {
                    begin:begin,
                    num:count
                },
            }).then(function successCallback(response) {
                // 请求成功执行代码
                dataList = response.data.ret; //这里需要修改
                //console.log('response.data = ',response.data);   
                // hide progress
                $scope.progress = { Ready: true };
                $scope.contactList = response.data.ret;
                //return dataList;
            }, function errorCallback(response) {
            // 请求失败执行代码
                alert("Http Get Error:GetTNovelTableInfo");

            });
        }
        // init socialite jquery plugin
        Socialite.load();
    }]);
