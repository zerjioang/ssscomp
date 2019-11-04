from setuptools import setup, find_packages

__version__ = "0.1"

install_requires = [
    'mkdocs', 'mkdocs-material'
]

setup(
    name="ssscomp-docs",
    version=__version__,
    packages=find_packages(),
    install_requires=install_requires,
    package_data={
        '': ['*.rst'],
    },
    author="sergio.anguita",
    author_email="sergio.anguita@opendeusto.es",
    description="Documentation files for ssscomp",
    license="GPL3",
    keywords="docker ssscomp docs",
    url="https://github.com/zerjioang/ssscomp",
    classifiers=[
        'Development Status :: 4 - Beta',
        'Environment :: Other Environment',
        'Intended Audience :: Developers',
        'Operating System :: OS Independent',
        'Programming Language :: Python',
        'Programming Language :: Python :: 2.7',
        'Programming Language :: Python :: 3.4',
        'Programming Language :: Python :: 3.5',
        'Programming Language :: Python :: 3.6',
        'Topic :: Utilities',
    ],
)
