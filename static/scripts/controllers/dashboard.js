'use strict';

angular.module('gatorloopWebApp')
    .controller('dashboardCTRL', function($scope, dashboardService, $timeout) {
        $scope.isStreaming = false;
        $scope.currentVelocity = {x: 0, y: 0};
        $scope.velocities = [
            0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
            0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
            0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
            0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
            0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
            0, 0, 0, 0, 0, 0, 0, 0 , 0, 0
        ];
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
            //roll, pitch, yaw values
            r: 0,
            p: 0,
            y: 0
        };
        $scope.rotations = [];
        $scope.currentAcceleration = 0;
        $scope.accelerations = [
            0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
            0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
            0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
            0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
            0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
            0, 0, 0, 0, 0, 0, 0, 0 , 0, 0
        ];
        $scope.time = 0;
        $scope.percentage = 0;
        $scope.distanceLeft = 1609;


//////////////   Here's where the graph code starts /////////////////////////

        var limit = 60 * 1,
                duration = 750,
                now = new Date(Date.now() - duration);

        var width = 600,
                height = 325;

        var groups = {
                current: {
                    color: 'orange',
                    data: $scope.velocities
                },
                target: {
                    color: 'aqua',
                    data: $scope.accelerations
                }
            };

        var x = d3.time.scale()
                .domain([now - (limit - 2), now - duration])
                .range([0, width])

        var y = d3.scale.linear()
                .domain([0, 250])
                .range([height, 0])

        var line = d3.svg.line()
                .interpolate('basis')
                .x(function(d, i) {
                    return x(now - (limit - 1 - i) * duration)
                })
                .y(function(d) {
                    return y(d)
                })

        var svg = d3.select('.graph').append('svg')
                .attr('class', 'chart')
                .attr('width', width)
                .attr('height', height + 50)

        var axis = svg.append('g')
                .attr('class', 'x axis')
                .attr('transform', 'translate(0,' + height + ')')
                .call(x.axis = d3.svg.axis().scale(x).orient('bottom'))

        svg.append("g")
            .attr("class", "y axis")
            //.attr('transform', 'translate(0,' + height + ')')
            .call( y.axis = d3.svg.axis().scale(y).orient('right'))

        var paths = svg.append('g');

            for (var name in groups) {
                var group = groups[name];
                group.path = paths.append('path')
                    .data([group.data])
                    .attr('class', name + ' group')
                    .style('stroke', group.color)
            }

        function tick() {
                now = new Date();
                for (var name in groups) {
                    var group = groups[name];
                    if($scope.isStreaming===false) group.data.push(0);
                    group.path.attr('d', line);
                    //console.log("velocity: " + (20 + Math.random() * 100));
                    //console.log("acceleration: " + $scope.currentAcceleration);
                    //console.log("group: " + group.data)
                }

                // Shift domain
                x.domain([now - (limit - 2) * duration, now - duration]);

                // Slide x-axis left
                axis.transition()
                    .duration(duration)
                    .ease('linear')
                    .call(x.axis);

                // Slide paths left
                paths.attr('transform', null)
                    .transition()
                    .duration(duration)
                    .ease('linear')
                    .attr('transform', 'translate(' + x(now - (limit - 1) * duration) + ')')
                    .each('end', tick);

                // Remove oldest data point from each group
                for (var name in groups) {
                    var group = groups[name];
                    group.data.shift();
                }
        } tick();


      ////////////////// Here's where it stops /////////////////////////////////////////////////

      $scope.sendStopSignal = function() {
          dashboardService.sendStopSignal().success(function(data) {
              console.log(data.stop);
          }).error(function(data) {
              console.error("Error", data);
          });
      }

      $scope.getCurrentVelocity = function() {
          dashboardService.get("velocity").success(function(data) {
              $scope.currentVelocity = data.velocity;
              $scope.velocities.push(data.velocity);
          }).error(function(data) {
              console.error("Error", data);
          });
      }

      $scope.getCurrentTemperature = function() {
          dashboardService.get("temperature").success(function(data) {
              $scope.currentTemperature = data.temperature;
              $scope.temperatures.push(data.temperature);
          }).error(function(data) {
              console.error("Error", data);
          });
      }



      $scope.getCurrentRotations = function() {
          dashboardService.get("rotation").success(function(data) {
              $scope.currentRotation = {r: data.roll, p: data.pitch, y: data.yaw};
              $scope.rotations.push({r: data.roll, p: data.pitch, y: data.yaw});

          }).error(function(data) {
              console.error("Error", data);
          });
      }

      $scope.getCurrentAcceleration = function() {
          dashboardService.get("acceleration").success(function(data) {
              $scope.currentAcceleration = data.acceleration;
              $scope.accelerations.push(data.acceleration);
          }).error(function(data) {
              console.error("Error", data);
          });
      }

      $scope.getCurrentPrimaryBattery = function() {
        dashboardService.get("primarybattery").success(function(data) {
            $scope.currentPrimaryBattery = {vol: data.voltage, soc: data.soc, tmp: data.temperature, amp: data.amp_hours };
            $scope.primaryBatterys.push({vol: data.vol, soc: data.soc, tmp: data.tmp, amp: data.amp_hours });
        }).error(function(data) {
            console.error("Error", data);
        });
      }

      $scope.getCurrentAuxiliaryBattery = function() {
        dashboardService.get("auxbattery").success(function(data) {
            $scope.currentAuxiliaryBattery = {vol: data.voltage, soc: data.soc, tmp: data.temperature, amp: data.amp_hours };
            $scope.auxiliaryBatterys.push({vol: data.vol, soc: data.soc, tmp: data.tmp, amp: data.amp_hours });
        }).error(function(data) {
            console.error("Error", data);
        });
      }

      $scope.startGettingData = function(){
        $scope.interval = setInterval(function() {
                $scope.isStreaming = true;
                $scope.getCurrentVelocity();
                $scope.getCurrentAcceleration();
                $scope.getCurrentRotations();
                $scope.getCurrentTemperature();
                $scope.getCurrentPosition();
                $scope.getCurrentPrimaryBattery();
                $scope.getCurrentAuxiliaryBattery();
                $scope.setSecondaryBatteryLevel()
                $scope.setPrimaryBatteryLevel();
                $scope.setSecondaryBatteryLevel();
                $scope.setPodPosition();
                $scope.setDistanceLeft();
                //tick();//$scope.graph();
              }, 750);
      }

      $scope.stopGettingData = function() {
          $scope.isStreaming = false;
          clearInterval($scope.interval);
      }

      $scope.setPrimaryBatteryLevel = function() {

        dashboardService.get("primarybattery").success(function(data){ $scope.percentage = 100*(data.voltage / 10); });
          document.getElementById("primaryBatteryLevel").style.top = $scope.percentage + "%";
          document.getElementById("primaryBatteryLevel").style.height = 100 - $scope.percentage + "%";
          console.log($scope.percentage);
      }

      $scope.setSecondaryBatteryLevel = function() {
          dashboardService.get("auxbattery").success(function(data){ $scope.percentage = 100*(data.voltage / 10); });
          document.getElementById("secondaryBatteryLevel").style.top = $scope.percentage + "%";
          document.getElementById("secondaryBatteryLevel").style.height = 100 - $scope.percentage + "%";
          /*if($scope.percentage < 60) {
              document.getElementById("secondaryBatteryLevel").style.background = ""
          }*/
          console.log($scope.percentage);
      }

      $scope.setSecondaryBatteryLevel = function() {
          dashboardService.get("auxbattery").success(function(data){ $scope.percentage = 100*(data.voltage / 10); });
          document.getElementById("secondaryBatteryLevel").style.top = $scope.percentage + "%";
          document.getElementById("secondaryBatteryLevel").style.height = 100 - $scope.percentage + "%";
            /*if($scope.percentage < 60) {
             document.getElementById("secondaryBatteryLevel").style.background = ""
             }*/
          console.log($scope.percentage);
      }

      $scope.setDistanceLeft = function() {
            $scope.distanceLeft = 1609 - $scope.currentPosition;
            if($scope.distanceLeft < 0) $scope.distanceLeft = 0;
      }

      $scope.getCurrentPosition = function() {
            dashboardService.get("position").success(function(data) {
                $scope.currentPosition = data.position;
                $scope.positions.push(data.position);
            }).error(function(data) {
                console.error("Error", data);
            });
      }

      $scope.setPodPosition = function() {
          if($scope.distanceLeft = 0) $scope.currentPosition = 1609;
          document.getElementById("pod_move_control").style.marginLeft = 9 + 336*($scope.currentPosition / 1609) + "%";
      }
        
    });
