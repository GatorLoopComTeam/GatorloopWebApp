'use strict';

angular.module('gatorloopWebApp')
    .controller('dashboardCTRL', function($scope, dashboardService, $timeout) {
      $scope.currentVelocity = {x: 0, y: 0};
      $scope.velocities = [];
      $scope.currentTemperature = 0;
      $scope.temperatures = [];
      $scope.currentPosition = 0;
      $scope.positions = [];
      $scope.positionPercentage = 0;
      $scope.primaryBatterys = [];
      $scope.currentPrimaryBattery = {
        vol: 0,
        soc: 0,
        tmp: 0,
        amp: 0
      }
      $scope.auxiliaryBatterys = [];
      $scope.currentAuxiliaryBattery = {
        vol: 0,
        soc: 0,
        tmp: 0,
        amp: 0
      }
      $scope.currentRotation = {
        r: 0,
        p: 0,
        y: 0
      };
      $scope.rotations = [];
      $scope.currentAcceleration = 0;
      $scope.accelerations = [];
      $scope.time = 0;


      $scope.sendStopSignal = function() {
          dashboardService.sendStopSignal().success(function(data) {
              console.log(data.stop);
          }).error(function(data) {
              console.error("Error", data);
          });
      };

      $scope.getCurrentVelocity = function() {
          dashboardService.get("velocity").success(function(data) {
              $scope.currentVelocity = data.velocity;
              $scope.velocities.push(data.velocity);
          }).error(function(data) {
              console.error("Error", data);
          });
      };

      $scope.getCurrentTemperature = function() {
          dashboardService.get("temperature").success(function(data) {
              $scope.currentTemperature = data.temperature;
              $scope.temperatures.push(data.temperature);
          }).error(function(data) {
              console.error("Error", data);
          });
      };

      $scope.getCurrentPosition = function() {
          dashboardService.get("position").success(function(data) {
              $scope.currentPosition = data.position;
              $scope.positions.push(data.position);
          }).error(function(data) {
              console.error("Error", data);
          });
      };

      $scope.getCurrentRotations = function() {
          dashboardService.get("rotation").success(function(data) {
              $scope.currentRotation = {r: data.roll, p: data.pitch, y: data.yaw};
              $scope.rotations.push({r: data.roll, p: data.pitch, y: data.yaw});

          }).error(function(data) {
              console.error("Error", data);
          });
      };

      $scope.getCurrentAcceleration = function() {
          dashboardService.get("acceleration").success(function(data) {
              $scope.currentAcceleration = data.acceleration;
              $scope.accelerations.push(data.acceleration);
          }).error(function(data) {
              console.error("Error", data);
          });
      };

      $scope.getCurrentPrimaryBattery = function() {
        dashboardService.get("primarybattery").success(function(data) {
            $scope.currentPrimaryBattery = data;
            $scope.primaryBatterys.push(data);
        }).error(function(data) {
            console.error("Error", data);
        });
      }

      $scope.getCurrentAuxiliaryBattery = function() {
        dashboardService.get("auxbattery").success(function(data) {
            $scope.currentAuxiliaryBattery = data;
            $scope.auxiliaryBatterys.push(data);
        }).error(function(data) {
            console.error("Error", data);
        });
      }

      $scope.startGettingData = function(){
        $scope.interval = setInterval(function() {
                $scope.getCurrentVelocity();
                $scope.getCurrentAcceleration();
                $scope.getCurrentRotations();
                $scope.getCurrentTemperature();
                $scope.getCurrentPosition();
                $scope.getCurrentPrimaryBattery();
                $scope.getCurrentAuxiliaryBattery();
              }, 1000);
      }

      $scope.stopGettingData = function() {
        clearInterval($scope.interval);
      }

      $scope.setPrimaryBatteryLevel = function(percent) {
        document.getElementById("primaryBatteryLevel").style.height = (percent/100)*85 + "%";
      }

      $scope.setSecondaryBatteryLevel = function(percent) {
        document.getElementById("secondaryBatteryLevel").style.height = (percent/100)*85 + "%";
      }

      // var level = 2;
      // var levelInt = setInterval(function() {
      //   document.getElementById("podIconPusher").style.width = level + "%";
      //
      //   level += 0.2;
      //   if(level > 92) clearInterval(levelInt);
      //   // $scope.setPrimaryBatteryLevel(level);
      //   // $scope.setSecondaryBatteryLevel(level);
      // }, 10);

    });
