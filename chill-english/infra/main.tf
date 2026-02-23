locals {
  project_id = "chill-english-dev"
  region     = "asia-northeast1"
}

terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 6.0"
    }
    google-beta = {
      source  = "hashicorp/google-beta"
      version = "~> 6.0"
    }
  }
}

provider "google" {
  project               = local.project_id
  region                = local.region
  user_project_override = true
  billing_project       = local.project_id
}

provider "google-beta" {
  project               = local.project_id
  region                = local.region
  user_project_override = true
  billing_project       = local.project_id
}

# API 有効化
resource "google_project_service" "tts" {
  service            = "texttospeech.googleapis.com"
  disable_on_destroy = false
}

resource "google_project_service" "firestore" {
  service            = "firestore.googleapis.com"
  disable_on_destroy = false
}

resource "google_project_service" "firebase" {
  provider           = google-beta
  service            = "firebase.googleapis.com"
  disable_on_destroy = false
}

# API キー管理 API 有効化
resource "google_project_service" "apikeys" {
  service            = "apikeys.googleapis.com"
  disable_on_destroy = false
}

# Identity Platform API 有効化（Authentication に必要）
resource "google_project_service" "identitytoolkit" {
  service            = "identitytoolkit.googleapis.com"
  disable_on_destroy = false
}

# TTS 用 API キー
resource "google_apikeys_key" "tts" {
  name         = "chill-english-tts"
  display_name = "Chill English TTS"

  restrictions {
    api_targets {
      service = "texttospeech.googleapis.com"
    }
  }

  depends_on = [google_project_service.tts, google_project_service.apikeys]
}

# Firebase プロジェクト
resource "google_firebase_project" "default" {
  provider = google-beta
  depends_on = [google_project_service.firebase]
}

# 匿名認証
resource "google_identity_platform_config" "default" {
  provider = google-beta

  sign_in {
    anonymous {
      enabled = true
    }
  }

  depends_on = [google_project_service.identitytoolkit, google_firebase_project.default]
}

# Firestore データベース
resource "google_firestore_database" "default" {
  name        = "(default)"
  location_id = local.region
  type        = "FIRESTORE_NATIVE"

  depends_on = [google_firebase_project.default]
}

# Firebase Web App
resource "google_firebase_web_app" "default" {
  provider     = google-beta
  display_name = "Chill English"

  depends_on = [google_firebase_project.default]
}

data "google_firebase_web_app_config" "default" {
  provider   = google-beta
  web_app_id = google_firebase_web_app.default.app_id
}
