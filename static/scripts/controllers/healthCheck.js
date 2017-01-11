'use strict';

angular.module('gatorloopWebApp')
    .controller('healthCheckCTRL', function($scope, dashboardService) {
      $scope.batteryTempLowerBound = -20;
      $scope.batteryTempUpperBound = 70;
      $scope.batteryVoltageLowerBound = 32;
      $scope.batteryVoltageUpperBound = 41.4;
      $scope.batteryAHLowerBound = 8;
      $scope.batteryAHUpperBound = 40;
      $scope.batterySOCLowerBound = 60;
      $scope.batterySOCUpperBound = 100;

      $scope.currentState = "";
      $scope.primaryBrakesEngaged = false;
      $scope.auxiliaryBrakesEngaged = false;

      $scope.getPrimaryBattery = function() {
        dashboardService.get("primarybattery").success(function(data) {
          console.log("primaryBattery");
          console.log(data);
          $scope.primaryBattery = data;
        }).error(function(data) {
            console.error("Error", data);
        });
      };

      $scope.getAuxiliaryBattery = function() {
        dashboardService.get("auxbattery").success(function(data) {
          console.log("auxBattery");
          console.log(data);
          $scope.auxBattery = data;
        }).error(function(data) {
            console.error("Error", data);
        });
      };

      $scope.getCurrentState = function() {
        dashboardService.get("state").success(function(data) {
          console.log("state");
          console.log(data);
          $scope.currentState = data.state;
        }).error(function(data) {
          console.error("Error: ", data);
        });
      };

      $scope.getBrakeStatus = function() {
        dashboardService.get("brakestatus").success(function(data) {
          console.log("brakestatus");
          console.log(data);
          $scope.primaryBrakesEngaged = data.primary_engaged;
          $scope.auxiliaryBrakesEngaged = data.auxiliary_engaged;
        })
      }

      $scope.startGettingData = function(){
        $scope.interval = setInterval(function() {
                $scope.getPrimaryBattery();
                $scope.getAuxiliaryBattery();
                $scope.getCurrentState();
                $scope.getBrakeStatus();
              }, 2000);
      }

      $scope.startGettingData();
    });
