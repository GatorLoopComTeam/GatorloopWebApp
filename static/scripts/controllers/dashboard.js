'use strict';

angular.module('gatorloopWebApp')
    .controller('dashboardCTRL', function($scope, dashboardService, $timeout) {
        $scope.isStreaming = false;

        $scope.percentage = 0;
        $scope.distanceLeft = 1609;


      $scope.emergencyStop = function() {
          dashboardService.emergencyStop().success(function(data) {
              alert("POD STOPPED");
          }).error(function(data) {
              alert("NOT STOPPED");
          });
      }

      $scope.killPod = function() {
        dashboardService.killPod().success(function(data) {
          alert("POD KILLED");
        }).error(function(data) {
          alert("NOT KILLED");
        });
      }

      $scope.getCurrentVelocity = function() {
          dashboardService.get("velocity").success(function(data) {
              $scope.currentVelocity = data.velocity;
          }).error(function(data) {
              console.error("Error", data);
          });
      }

      $scope.getCurrentAcceleration = function() {
          dashboardService.get("acceleration").success(function(data) {
              $scope.currentAcceleration = data.acceleration;
          }).error(function(data) {
              console.error("Error", data);
          });
      }

      $scope.getCurrentPrimaryBattery = function() {
        dashboardService.get("primarybattery").success(function(data) {
            $scope.currentPrimaryBattery = {vol: data.voltage, soc: data.soc, tmp: Math.max(data.pack1_temp, data.pack2_temp, data.pack3_temp), amp: data.amp_hours };
        }).error(function(data) {
            console.error("Error", data);
        });
      }

      $scope.getCurrentAuxiliaryBattery = function() {
        dashboardService.get("auxbattery").success(function(data) {
            $scope.currentAuxiliaryBattery = {vol: data.voltage, soc: data.soc, tmp: Math.max(data.pack1_temp, data.pack2_temp, data.pack3_temp), amp: data.amp_hours };
        }).error(function(data) {
            console.error("Error", data);
        });
      }

      $scope.getCurrentState = function() {
        dashboardService.get("state").success(function(data) {
          $scope.currentState = data.state;
        }).error(function(data) {
          console.error("Error: ", data);
        });
      };

      $scope.getPodHealthy = function() {
        $scope.healthy = $scope.currentPrimaryBattery.tmp < 70 &&
        $scope.currentPrimaryBattery.tmp > -20 &&
        $scope.currentAuxiliaryBattery.tmp < 70 &&
        $scope.currentAuxiliaryBattery.tmp > -20 &&
        $scope.currentPrimaryBattery.vol > 32 &&
        $scope.currentPrimaryBattery.vol < 40.8 &&
        $scope.currentAuxiliaryBattery.vol > 32 &&
        $scope.currentAuxiliaryBattery.vol < 40.8 &&
        $scope.currentPrimaryBattery.soc > 60 &&
        $scope.currentPrimaryBattery.soc < 100 &&
        $scope.currentAuxiliaryBattery.soc > 60 &&
        $scope.currentAuxiliaryBattery.soc < 100 &&
        $scope.currentPrimaryBattery.amp > 8 &&
        $scope.currentPrimaryBattery.amp < 40 &&
        $scope.currentAuxiliaryBattery.amp > 8 &&
        $scope.currentAuxiliaryBattery.amp < 40;
      }

      $scope.startGettingData = function(){
        $scope.interval = setInterval(function() {
                $scope.getCurrentVelocity();
                $scope.getCurrentAcceleration();
                $scope.getCurrentPosition();
                $scope.getCurrentPrimaryBattery();
                $scope.getCurrentAuxiliaryBattery();
                $scope.setSecondaryBatteryLevel()
                $scope.setPrimaryBatteryLevel();
                $scope.setSecondaryBatteryLevel();
                $scope.getPodHealthy();
                $scope.getCurrentState();
              }, 200);
      }
      $scope.startGettingData();

      $scope.stopGettingData = function() {
          $scope.isStreaming = false;
          clearInterval($scope.interval);
      }

      $scope.setPrimaryBatteryLevel = function() {
        console.log("soc = " + $scope.currentPrimaryBattery.soc);
          document.getElementById("primaryBatteryLevel").style.height = $scope.currentPrimaryBattery.soc + "%";
          document.getElementById("primaryBatteryLevel").style.top = 100 - $scope.currentPrimaryBattery.soc + "%";
      }

      $scope.setSecondaryBatteryLevel = function() {
        console.log("soc = " + $scope.currentAuxiliaryBattery.soc);
          document.getElementById("secondaryBatteryLevel").style.height = $scope.currentAuxiliaryBattery.soc + "%";
          document.getElementById("secondaryBatteryLevel").style.top = 100 - $scope.currentAuxiliaryBattery.soc + "%";
      }

      $scope.getCurrentPosition = function() {
            dashboardService.get("position").success(function(data) {
                if(data.position <= 1609) $scope.currentPosition = data.position;
                else $scope.currentPosition = 1609;
                $scope.positions.push(data.position);
                $scope.positions.push(1609);
            }).error(function(data) {
                console.error("Error", data);
            });
        }


    });
