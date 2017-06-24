"use strict";

/* controllers module */

angular.module("BsTableDirective.SysNoticeControllers", ["BsTableDirective.Services","ui.bootstrap"])
    .controller("SysNoticeCtrl", ["$scope", "BootswatchService","$http", 
        function ($scope, BootswatchService, $http) {
        // scope models
        $scope.page = { Css: "/statics/bootstrap/css/bootstrap.min.css" };
        $scope.showDay = [
            {name : "1天", type : 1},
            {name : "2天", type : 2},
            {name : "5天", type : 5},
            {name : "10天", type : 10},
        ];
        $scope.contact = "";
        $scope.selecttime = $scope.showDay[0];
        $scope.progress = { Ready: false };
        $scope.show = "";

        // get themes from bootswatch
        BootswatchService.GetAll().success(function (result) {
            // set themes to scope
            $scope.page.Themes = result.themes;
            $scope.page.Css = $scope.page.Themes[15].cssMin;
        });

        // get data for bs-table
        GenerateData();
        $scope.SaveNotice = function(){
            $http({
                method: 'GET',
                url: '/ChangeSystemNoticeJson',
                params: {
                        notice:$scope.contact,
                        time:$scope.selecttime.type
                    },

            }).then(function successCallback(response) {
                // 请求成功执行代码
                console.log('response.data= ', response.data); 
                if (response.data.code == 1) {
                    alert("保存成功");
                }else{
                    alert("保存失败:" + response.data.ret);
                }

            }, function errorCallback(response) {
            // 请求失败执行代码
                alert("Https Error:ChangeSystemNoticeJson");

            });
        }
        // generate data
        function GenerateData() {
            $http({
                method: 'GET',
                url: '/GetSystemNoticeJson'
            }).then(function successCallback(response) {
                // 请求成功执行代码
                console.log('GetSystemNotice ret :', response.data); 
                if (response.data.code == 1) {
                    $scope.contact = response.data.ret;
                    $scope.show = "当前公告如下";
                }else if(response.data.code == -1){
                    $scope.show = "Redis错误，请查看";
                    alert("redis error.pls check");
                }else if(response.data.code == 0){
                    $scope.show = "当前公告为空";
                    $scope.contact = "";
                }
                // hide progress
                $scope.progress = { Ready: true };
                //$scope.$apply();
            }, function errorCallback(response) {
            // 请求失败执行代码
                alert("Https Get Error:GetSystemNoticeJson");

            });
        }
        // init socialite jquery plugin
        Socialite.load();
    }]);
    