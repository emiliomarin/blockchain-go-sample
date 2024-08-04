const app = artifacts.require("./App");

module.exports = function(deployer) {
    deployer.deploy(app);
}