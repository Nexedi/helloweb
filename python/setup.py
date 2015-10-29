# minimal egg specification for helloweb.py
from setuptools import setup

setup(
    name    = 'helloweb',
    entry_points = { 'console_scripts': ['helloweb = helloweb:main'] },

    # NOTE we can require other eggs via
    # install_requires = [ ... ]
)
