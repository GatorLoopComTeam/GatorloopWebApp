'use strict';

angular.module('gatorloopWebApp')
    .controller('dashboardCTRL', function($scope, dashboardService) {

        $scope.newVelocity = 0.00;

      $scope.sendStopSignal = function() {
          dashboardService.sendStopSignal().success(function(data) {
              alert(data.stop);
          }).error(function(data) {
              alert("Error", data);
          });
      };

      $scope.getCurrentVelocity = function() {
          dashboardService.get("velocity").success(function(data) {
              alert(data.velocity);
          }).error(function(data) {
              alert("Error", data);
          });
      };

      $scope.getCurrentPressure = function() {
          dashboardService.get("pressure").success(function(data) {
              alert(data.pressure);
          }).error(function(data) {
              alert("Error", data);
          });
      };

      $scope.getCurrentTemperature = function() {
          dashboardService.get("temperature").success(function(data) {
              alert(data.temperature);
          }).error(function(data) {
              alert("Error", data);
          });
      };

      $scope.getCurrentPosition = function() {
          dashboardService.get("position").success(function(data) {
              alert(data.position);
          }).error(function(data) {
              alert("Error", data);
          });
      };

      $scope.getCurrentRotations = function() {
          dashboardService.get("rotation").success(function(data) {
              alert("roll = " + data.roll + ", pitch = " + data.pitch + ", yaw = " + data.yaw);
          }).error(function(data) {
              alert("Error", data);
          });
      };

      $scope.getCurrentAcceleration = function() {
          dashboardService.get("acceleration").success(function(data) {
              alert(data.acceleration);
          }).error(function(data) {
              alert("Error", data);
          });
      };


    });
