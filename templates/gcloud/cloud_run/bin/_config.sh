export PYTHON_VERSION=3.7.0

export VIRTUALENV_NAME=$(pyenv local) || ""
export REPO_ROOT=$(cd $(dirname $0)/.. && pwd)
