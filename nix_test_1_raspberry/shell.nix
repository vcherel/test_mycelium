{ nixpkgs ? import <nixpkgs> {} }:

let
  myQEMU = nixpkgs.qemu.override {
  };
in

nixpkgs.mkShell {
  buildInputs = [
    myQEMU
    # Add any other dependencies you need here
  ];

  shellHook = ''
    export LAUNCH="./launch.sh"
  '';

  meta = with nixpkgs.lib; {
    description = "QEMU environment with custom options";
    license = licenses.mit;
  };
}

