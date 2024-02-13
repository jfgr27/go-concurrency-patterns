{
  description = "Go concurrency pattern";

  inputs.flake-utils.url = "github:numtide/flake-utils";

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        p = nixpkgs.legacyPackages.${system};
      in
        {
          devShells = rec {
            default = p.mkShell {
              packages = [];
            };
          };

        }
    );
}
