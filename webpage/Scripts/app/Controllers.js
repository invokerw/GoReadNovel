"use strict";

/* controllers module */

angular.module("BsTableDirective.Controllers", ["BsTableDirective.Services","ui.bootstrap"])
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
    
    .controller("NovelInfoCtrl", ["$scope", "$uibModal","BootswatchService", "$http", function ($scope,$uibModal, BootswatchService,$http) {
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
            //alert(JSON.stringify(contact));
            $scope.Open(contact,true);
        };
        // edit function for bs-table
        $scope.Edit = function (contact) {
            //alert(JSON.stringify(contact));
            $scope.Open(contact,false);
        };
        // remove function for bs-table
        $scope.Remove = function (contact) {
            //alert(JSON.stringify(contact));
            if (confirm("你确定删除该项？")) {  
                //alert("点击了确定");  
                $http({
                    method: 'GET',
                    url: '/GetDeleteNovleID',
                    params: {
                        novelid:contact.id
                    },
                }).then(function successCallback(response) {
                    // 请求成功执行代码
                    if(response.data.code == 1){
                        //删除成功
                        console.log('response.data = ',response.data);  
                        GenerateData(0,100);
                    } 
                 }, function errorCallback(response) {
                    // 请求失败执行代码
                    alert("Https Get Error:GetDeleteNovleID");

                });
            }  
            else {  
                //alert("闲得慌");  
                console.log('闲得慌啊你。。');  
             }  
        };
        
        var $ctrl = this;
        $ctrl.items = ['item1', 'item2', 'item3'];
        $scope.selected = ""
        $scope.Open = function(contact,readonly,size){
            //console.log("this.........open");
            var parentElem = undefined;
            //parentSelector ? angular.element($document[0].querySelector('.modal-demo ' + parentSelector)) : undefined;
            var modalInstance = $uibModal.open({
              animation: true,
              ariaLabelledBy: 'modal-title',
              ariaDescribedBy: 'modal-body',
              templateUrl: 'myModalContent.html',
              controller: 'ModalInstanceCtrl',
              controllerAs: '$ctrl',
              size: size,
              appendTo: parentElem,
              resolve: {
                items: function () {
                    return $ctrl.items;
                },
                contact:contact,
                readonly:readonly
              }
            });

            modalInstance.result.then(function (selectedItem) {
              $scope.selected = selectedItem;
            }, function () {
              console.log('Modal dismissed at: ' + new Date());
            });
        };
        
        // generate data
        function GenerateData(begin,count) {
            var dataList = [];
            $http({
                method: 'GET',
                url: '/GetNovelTableInfoJson',
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
    }])
    
    // Please note that $uibModalInstance represents a modal window (instance) dependency.
    // It is not the same as the $uibModal service used above.
    .controller('ModalInstanceCtrl',["$uibModalInstance","contact", "items","readonly","$http", function ($uibModalInstance, contact, items, readonly,$http) {
      var $ctrl = this;
      $ctrl.items = items;
      $ctrl.contact = contact;
      $ctrl.readonly = readonly
      //console.log("$ctrl.readonly = ",$ctrl.readonly)

      $ctrl.selected = {
        item: $ctrl.items[0]
      };

      $ctrl.ok = function () {
        //readonly = true 是Show
        if ($ctrl.readonly) {
            $uibModalInstance.close($ctrl.selected.item);
        }
        // false 是保存 Edit
        else{
            console.log("$ctrl.readonly false")
            $http({
                method: 'GET',
                url: '/GetEditNovelJson',
                params: {novel:$ctrl.contact},
            }).then(function successCallback(response) {
                // 请求成功执行代码

                if (response.data.code == 1) {
                    console.log("edit ok",response.data);
                    $uibModalInstance.close($ctrl.selected.item);
                }
                else {
                     console.log("edit err",response.data);
                     alert(response.data);
                }

            }, function errorCallback(response) {
            // 请求失败执行代码
                alert("Http Get Error:GetTNovelTableInfo");
            });

        }

      };

      $ctrl.cancel = function () {
        $uibModalInstance.dismiss('cancel');
      };
    }]);