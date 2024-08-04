// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

contract App {
    uint public sensorIndex;

    struct Sensor {
        uint id;
        string name;
        uint value;
        address user;
    }

    Sensor[] public sensors;
    mapping (uint => Sensor) public sensorByID; // maps sensor to ID
    mapping (address => Sensor[]) public userSensors; // maps user address to array of Sensors

    // Create sensor and set value
    // - external because it's called only from outside
    function createSensor(string calldata name, uint value) external returns(Sensor memory) {
        sensorIndex++;
        // memory because the variable only exists in this function
        Sensor memory sensor = Sensor({
            id: sensorIndex,
            name: name,
            value: value,
            user: msg.sender
        });

        sensors.push(sensor);
        userSensors[msg.sender].push(sensor);
        sensorByID[sensorIndex] = sensor;

        return sensor;
    }

    // get sensors linked to address making request
    // - external because it's called only from the outside
    // - view because it does not change any state
    function getSensors(address user) external view returns (Sensor[] memory) {
        return userSensors[user];
    }

    // Update sensor value. Returns an error if not owner.
    // - external because it's called only from outside
    function updateSensor(uint id, uint value) external returns(Sensor memory) {
        require(sensorByID[id].user == msg.sender, "sender is not owner"); // sender must be owner
        sensorByID[id].value = value;
        return sensorByID[id];
    }

    // get sensors linked to address making request
    // - public because it can be called externally and internally
    // - view because it does not change any state
    function getSensor(uint id) public view returns (Sensor memory) {
        return sensorByID[id];
    }
}