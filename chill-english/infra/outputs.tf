output "google_tts_api_key" {
  description = "Cloud TTS API キー (.env.local の GOOGLE_TTS_API_KEY に設定)"
  value       = google_apikeys_key.tts.key_string
  sensitive   = true
}

output "firebase_project_id" {
  description = "NEXT_PUBLIC_FIREBASE_PROJECT_ID"
  value       = local.project_id
}

output "firebase_auth_domain" {
  description = "NEXT_PUBLIC_FIREBASE_AUTH_DOMAIN"
  value       = data.google_firebase_web_app_config.default.auth_domain
}

output "firebase_storage_bucket" {
  description = "NEXT_PUBLIC_FIREBASE_STORAGE_BUCKET"
  value       = data.google_firebase_web_app_config.default.storage_bucket
}

output "firebase_api_key" {
  description = "NEXT_PUBLIC_FIREBASE_API_KEY"
  value       = data.google_firebase_web_app_config.default.api_key
  sensitive   = true
}

output "firebase_app_id" {
  description = "NEXT_PUBLIC_FIREBASE_APP_ID"
  value       = google_firebase_web_app.default.app_id
}

output "firebase_messaging_sender_id" {
  description = "NEXT_PUBLIC_FIREBASE_MESSAGING_SENDER_ID"
  value       = data.google_firebase_web_app_config.default.messaging_sender_id
}
