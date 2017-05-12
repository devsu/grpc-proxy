const Condor = require('condor-framework');

const Greeter = class {
  sayHello(call) {
    return { 'greeting': `Insecure: Hello ${call.request.name}`};
  }
};

const options = {
  'host': 'localhost',
  'port': 3001,
};

const server = new Condor(options);
server.addService('./hello.proto', 'myinsecureapp.Greeter', new Greeter());
server.start();