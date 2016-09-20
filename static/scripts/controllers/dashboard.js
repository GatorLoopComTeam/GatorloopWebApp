'use strict';

angular.module('gatorloopWebApp')
    .controller('dashboardCTRL', function($scope, dashboardService) {
      $scope.currentVelocity = 0;
      $scope.velocities = [];
      $scope.currentPressure = 0;
      $scope.pressures = [];
      $scope.currentTemperature = 0;
      $scope.temperatures = [];
      $scope.currentTemperature = 0;
      $scope.temperatures = [];
      $scope.currentPosition = 0;
      $scope.positions = [];
      $scope.currentRotation = {
        r: 0,
        p: 0,
        y: 0
      };
      var rotationObj = {
        r: 0,
        p: 0,
        y: 0
      };
      $scope.rotations = [];
      $scope.currentAcceleration = 0;
      $scope.accelerations = [];

      $scope.sendStopSignal = function() {
          dashboardService.sendStopSignal().success(function(data) {
              alert(data.stop);
          }).error(function(data) {
              alert("Error", data);
          });
      };

      $scope.getCurrentVelocity = function() {
          dashboardService.get("velocity").success(function(data) {
              $scope.currentVelocity = data.velocity;
              $scope.velocities.push(data.velocity);
          }).error(function(data) {
              alert("Error", data);
          });
      };

      $scope.getCurrentPressure = function() {
          dashboardService.get("pressure").success(function(data) {
              $scope.currentPressure = data.pressure;
              $scope.pressures.push(data.pressure);
          }).error(function(data) {
              alert("Error", data);
          });
      };

      $scope.getCurrentTemperature = function() {
          dashboardService.get("temperature").success(function(data) {
              $scope.currentTemperature = data.temperature;
              $scope.temperatures.push(data.temperature);
          }).error(function(data) {
              alert("Error", data);
          });
      };

      $scope.getCurrentPosition = function() {
          dashboardService.get("position").success(function(data) {
              $scope.currentPosition = data.position;
              $scope.positions.push(data.position);
          }).error(function(data) {
              alert("Error", data);
          });
      };

      $scope.getCurrentRotations = function() {
          dashboardService.get("rotation").success(function(data) {
              rotationObj.r = data.roll;
              rotationObj.p = data.pitch;
              rotationObj.y = data.yaw;
              $scope.currentRotation = rotationObj;
              $scope.rotations.push({r: data.roll, p: data.pitch, y: data.yaw});

          }).error(function(data) {
              alert("Error", data);
          });
      };

      $scope.getCurrentAcceleration = function() {
          dashboardService.get("acceleration").success(function(data) {
              $scope.currentAcceleration = data.acceleration;
              $scope.accelerations.push(data.acceleration);
          }).error(function(data) {
              alert("Error", data);
          });
      };

      setInterval(function() {
        $scope.getCurrentVelocity();
        $scope.getCurrentAcceleration();
        $scope.getCurrentRotations();
        $scope.getCurrentPressure();
        $scope.getCurrentTemperature();
        $scope.getCurrentPosition();
      }, 1000);

    });
