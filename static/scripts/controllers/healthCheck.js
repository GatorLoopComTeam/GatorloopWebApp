'use strict';

angular.module('gatorloopWebApp')
    .controller('healthCheckCTRL', function($scope, dashboardService) {
      $scope.batteryTempLowerBound = -20;
      $scope.batteryTempUpperBound = 70;
      $scope.batteryVoltageLowerBound = 32;
      $scope.batteryVoltageUpperBound = 40.8;
      $scope.batteryAHLowerBound = 8;
      $scope.batteryAHUpperBound = 40;
      $scope.batterySOCLowerBound = 60;
      $scope.batterySOCUpperBound = 100;
      $scope.currentProcTemp = 0;

      $scope.currentState = "";
      $scope.primaryBrakesEngaged = false;
      $scope.auxiliaryBrakesEngaged = false;
      $scope.currentVelocity = 0;
      $scope.currentAcceleration = 0;

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

      // $scope.getBrakeStatus = function() {
      //   dashboardService.get("brakestatus").success(function(data) {
      //     console.log("brakestatus");
      //     console.log(data);
      //     $scope.primaryBrakesEngaged = data.primary_engaged;
      //     $scope.auxiliaryBrakesEngaged = data.auxiliary_engaged;
      //   })
      // }

      $scope.sendEmergencyBrake = function() {
        dashboardService.get("/stop").success(function(data) {
          console.log("sent emergency brake");
          console.log(data);

          if (data.stop === true) {
            $scope.primaryBrakesEngaged = data.primary_engaged;
            $scope.auxiliaryBrakesEngaged = data.auxiliary_engaged;
          }
        }).error(function(err) {
          console.log("error engaging ebrake");
          console.log(err);
        });
      }

      $scope.sendKillPower = function() {
        dashboardService.get("/killpower").success(function(data) {
          console.log("sent kill power");
          console.log(data);

          if (data.kill_power === true) {
            $scope.powerOn = false;
          }
        }).error(function(err) {
          console.log("error killing power");
          console.log(err);
        });
      }

      $scope.getCurrentVelocity = function() {
          dashboardService.get("velocity").success(function(data) {
              $scope.currentVelocity = data.velocity;
          }).error(function(data) {
              console.error("Error", data);
          });
      }

      $scope.getCurrentCalcAcceleration = function() {
          dashboardService.get("calcAcceleration").success(function(data) {
              $scope.currentAcceleration = data.acceleration;
          }).error(function(data) {
              console.error("Error", data);
          });
      }

      $scope.getCurrentProcTemp = function() {
          dashboardService.get("procTemp").success(function(data) {
              $scope.currentProcTemp = data.temperature;
          }).error(function(data) {
              console.error("Error", data);
          });
      }

      $scope.sendReady = function() {
        dashboardService.get("ready").success(function(data) {
          if (data.ready === true) {
            console.log("ready recieved");
          } else {
            console.error("ready failed");
          }
        }).error(function(data) {
            console.error("Error", data);
        });
      }

      $scope.startGettingData = function(){
        $scope.interval = setInterval(function() {
                $scope.getPrimaryBattery();
                $scope.getAuxiliaryBattery();
                $scope.getCurrentState();
                $scope.getCurrentVelocity();
                $scope.getCurrentCalcAcceleration();
                $scope.getCurrentProcTemp();
              }, 200);
      }

      $scope.startGettingData();
    });
