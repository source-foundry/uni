#!/usr/bin/env python
# -*- coding: utf-8 -*-


import sys
import subprocess
import pytest


# Command line version, help, usage flag tests

def test_uni_main_version_request_short():
    out = subprocess.check_output(
        "./uni -v",
        stderr=subprocess.STDOUT,
        shell=True)

    assert "uni v" in str(out)
    assert "Unicode Standard v" in str(out)


def test_uni_main_version_request_long():
    out = subprocess.check_output(
        "./uni --version",
        stderr=subprocess.STDOUT,
        shell=True)

    assert "uni v" in str(out)
    assert "Unicode Standard v" in str(out)


def test_uni_main_usage_request():
    out = subprocess.check_output(
        "./uni --usage",
        stderr=subprocess.STDOUT,
        shell=True)

    assert "Usage:" in str(out)


def test_uni_main_help_request_short():
    out = subprocess.check_output(
        "./uni -h",
        stderr=subprocess.STDOUT,
        shell=True)

    assert "==========" in str(out)


def test_uni_main_help_request_long():
    out = subprocess.check_output(
        "./uni --help",
        stderr=subprocess.STDOUT,
        shell=True)

    assert "==========" in str(out)


# Standard input tests


def test_uni_main_stdin_search_glyphsearch_single():
    out = subprocess.check_output(
        "echo -n 'j' | ./uni",
        stderr=subprocess.STDOUT,
        shell=True)

    assert "U+006A 'j'" in str(out)


def test_uni_main_stdin_search_glyphsearch_multiple():
    out = subprocess.check_output(
        "echo -n 'ji' | ./uni",
        stderr=subprocess.STDOUT,
        shell=True)

    assert "U+006A 'j'" in str(out)
    assert "U+0069 'i'" in str(out)


def test_uni_main_stdin_search_unicodesearch_single():
    out = subprocess.check_output(
        """echo 006A | ./uni -g""",
        stderr=subprocess.STDOUT,
        shell=True)

    assert "U+006A 'j'" in str(out)


def test_uni_main_stdin_search_unicodesearch_multiple():
    out = subprocess.check_output(
        """echo 006A 0069 | ./uni -g""",
        stderr=subprocess.STDOUT,
        shell=True)

    assert "U+006A 'j'" in str(out)
    assert "U+0069 'i'" in str(out)


def test_uni_main_stdin_search_glyphsearch_empty_stdin():
    with pytest.raises(subprocess.CalledProcessError):
        out = subprocess.check_output(
            "cat /dev/null | ./uni",
            stderr=subprocess.STDOUT,
            shell=True)


def test_uni_main_stdin_search_glyphsearch_invalid_unicode():
    with pytest.raises(subprocess.CalledProcessError):
        out = subprocess.check_output(
            "echo ZZZZ | ./uni -g",
            stderr=subprocess.STDOUT,
            shell=True)


# Command argument tests


def test_uni_main_arg_search_glyphsearch_single():
    out = subprocess.check_output(
        "./uni j",
        stderr=subprocess.STDOUT,
        shell=True)

    assert "U+006A 'j'" in str(out)


def test_uni_main_arg_search_glyphsearch_multiple():
    out = subprocess.check_output(
        "./uni ji",
        stderr=subprocess.STDOUT,
        shell=True)

    assert "U+006A 'j'" in str(out)
    assert "U+0069 'i'" in str(out)


def test_uni_main_arg_search_unicodesearch_single():
    out = subprocess.check_output(
        """./uni -g 006A""",
        stderr=subprocess.STDOUT,
        shell=True)

    assert "U+006A 'j'" in str(out)


def test_uni_main_arg_search_unicodesearch_multiple():
    out = subprocess.check_output(
        """./uni -g 006A 0069""",
        stderr=subprocess.STDOUT,
        shell=True)

    assert "U+006A 'j'" in str(out)
    assert "U+0069 'i'" in str(out)


def test_uni_main_arg_search_glyphsearch_invalid_unicode():
    with pytest.raises(subprocess.CalledProcessError):
        out = subprocess.check_output(
            "./uni -g ZZZZ",
            stderr=subprocess.STDOUT,
            shell=True)