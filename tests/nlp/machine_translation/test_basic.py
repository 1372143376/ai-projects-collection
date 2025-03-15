import pytest

class TestMachineTranslation:
    def test_basic_translation(self):
        """Test basic text translation"""
        from translate import Translator
        translator = Translator(to_lang="zh")
        translation = translator.translate("Hello, world!")
        assert "你好" in translation or "世界" in translation

    def test_language_detection(self):
        """Test source language detection"""
        from langdetect import detect
        lang = detect("This is a test sentence.")
        assert lang == "en"

    def test_translation_quality(self):
        """Test translation quality assessment"""
        from bleu import compute_bleu
        reference = ["这是一个测试句子"]
        candidate = ["这是一个测试"]
        score = compute_bleu(reference, candidate)
        assert score > 0.5
