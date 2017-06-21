"use strict";

/* controllers module */

angular.module("BsTableDirective.UserControllers", ["BsTableDirective.Services","ui.bootstrap"])
    .controller("UserCtrl", ["$scope","$uibModal", "BootswatchService","$http", 
        function ($scope, $uibModal,BootswatchService, $http) {
        // scope models
        $scope.page = { Css: "/statics/bootstrap/css/bootstrap.min.css" };
        $scope.contactList = [];
        $scope.progress = { Ready: false };
        // get themes from bootswatch
        BootswatchService.GetAll().success(function (result) {
            // set themes to scope
            $scope.page.Themes = result.themes;
            $scope.page.Css = $scope.page.Themes[15].cssMin;
            // hide progress
            //$scope.progress = { Ready: true };
        });
        // get data for bs-table
        GenerateData();
        // show function for bs-table
        $scope.Show = function (contact) {
            //alert(JSON.stringify(contact));
            OpenNewWindow(contact, true, 0);
        };
        // edit function for bs-table
        /*$scope.Edit = function (contact) {
            alert(JSON.stringify(contact));
        };
        // remove function for bs-table
        $scope.Remove = function (contact) {
            alert(JSON.stringify(contact));
        };*/
        // generate data
        function GenerateData() {
            $http({
                method: 'GET',
                url: '/GetUsersInfoJson'
            }).then(function successCallback(response) {
                // 请求成功执行代码
                console.log('response.data = ',response.data);   
                // hide progress
                $scope.progress = { Ready: true };
                $scope.contactList = response.data.ret;
                //return dataList;
            }, function errorCallback(response) {
            // 请求失败执行代码
                alert("Http Get Error:GetTNovelTableInfo");

            });
        }
        var $ctrl = this;

        function OpenNewWindow(contact,readonly,ty,size){
            //console.log("this.........open");
            var parentElem = undefined;
            //parentSelector ? angular.element($document[0].querySelector('.modal-demo ' + parentSelector)) : undefined;
            var modalInstance = $uibModal.open({
              animation: true,
              ariaLabelledBy: 'modal-title',
              ariaDescribedBy: 'modal-body',
              templateUrl: 'userWindow.html',
              controller: 'ShowUserInfoCtrl',
              controllerAs: '$ctrl',
              size: size,
              appendTo: parentElem,
              resolve: {
                contact:contact,
                opentype:ty,
                readonly:readonly
          
              }
            });
        };
        // init socialite jquery plugin
        Socialite.load();
    }])
    .controller('ShowUserInfoCtrl',["$uibModalInstance","contact", "readonly","$http","opentype",
     function ($uibModalInstance, contact,  readonly,$http,opentype) {
      var $ctrl = this;

      $ctrl.contact = contact;
      console.log($ctrl.contact);
      $ctrl.readonly = readonly;
      $ctrl.opentype = opentype;
      $ctrl.title = "User";
      //console.log("$ctrl.opentype = ",$ctrl.opentype)
      if ($ctrl.opentype == 0) {
        $ctrl.title = "UserInfo Detail";
      }else if($ctrl.opentype == 1){
        $ctrl.title = "Edit UserInfo";
      }else if($ctrl.opentype == 2){
        $ctrl.title = "New UserInfo";
      }

      $ctrl.ok = function () {
        //readonly = true 是Show
        if($ctrl.opentype == 0) {
            console.log("$ctrl.opentype 0")
        }
        // false 是保存 Edit
        else if($ctrl.opentype == 1) {
            console.log("$ctrl.opentype 1")
        }
        else if($ctrl.opentype == 2) {
            console.log("$ctrl.opentype 2");
        }
        $uibModalInstance.dismiss('cancel');
      };

      $ctrl.cancel = function () {
        $uibModalInstance.dismiss('cancel');
      };
    }]);
    ;
    