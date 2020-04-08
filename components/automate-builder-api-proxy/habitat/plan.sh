#shellcheck disable=SC2034
#shellcheck disable=SC2039
#shellcheck disable=SC2154

pkg_name="automate-builder-api-proxy"
pkg_origin="chef"
pkg_description="Wrapper package for habitat/builder-api-proxy"
pkg_version="0.1.0"
pkg_maintainer="Chef Software Inc. <support@chef.io>"
pkg_license=("Chef-MLSA")
pkg_upstream_url="https://www.chef.io/automate"
pkg_svc_user="root"
pkg_deps=(
  chef/mlsa
  # core/bash is pinned to a version that is compatible with the pinned version
  # of habitat/builder-api-proxy. When that pin is updated this pin should be removed.
  core/bash/4.4.19/20190115012619
  "${local_platform_tools_origin:-chef}/automate-platform-tools"
  # We need to pin here to get a build from unstable
  "habitat/builder-api-proxy/8796/20200309134228"
)

pkg_build_deps=(
  core/gcc
)

pkg_binds=(
  [automate-builder-api]="http-port"
)

pkg_exports=(
  [port]=service.port
)

pkg_exposes=(port)

pkg_bin_dirs=(bin)

pkg_scaffolding="${local_scaffolding_origin:-chef}/automate-scaffolding"

do_unpack() {
    return 0
}

do_build(){
    return 0
}

do_install() {
    mkdir -p "${pkg_prefix}/config/"
    proxypath="$(pkg_path_for habitat/builder-api-proxy)"
    proxyrel="${proxypath##*/}"
    sed -r "s/BUILDER_API_PROXY_RELEASE/${proxyrel}/g" support/index.html > "${pkg_prefix}/config/index.html"

    return 0
}

do_strip() {
    return 0
}
