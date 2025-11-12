use axum::{
    routing::get,
    Router,
};
use sqlx::MySqlPool;
use mongodb::{Client, options::ClientOptions};
use std::net::SocketAddr;

#[tokio::main]
async fn main() {

    tracing_subscriber::fmt::init();

    let maria_url = std::env::var("DATABASE_URL_MARIA")
        .expect("DATABASE_URL_MARIA must be set");
    let maria_pool = MySqlPool::connect(&maria_url).await
        .expect("MariaDB connection failed");
    tracing::info!("MariaDB Connected!");

    let mongo_url = std::env::var("DATABASE_URL_MONGO")
        .expect("DATABASE_URL_MONGO must be set");
    let mongo_options = ClientOptions::parse(&mongo_url).await
        .expect("MongoDB parse failed");
    let mongo_client = Client::with_options(mongo_options)
        .expect("MongoDB client creation failed");
    let _ = mongo_client.list_database_names(None, None).await
        .expect("MongoDB connection failed (ping)");
    tracing::info!("MongoDB Connected!");


    let app = Router::new()
        .route("/", get(handler_hello))
        .with_state((maria_pool, mongo_client)); 

    let addr = SocketAddr::from(([0, 0, 0, 0], 8080));
    
    let listener = tokio::net::TcpListener::bind(addr).await
        .unwrap();
    
    tracing::info!("listening on {}", listener.local_addr().unwrap());
    
    axum::serve(listener, app) // <-- axum::serve を使用
        .await
        .unwrap();
}

async fn handler_hello() -> &'static str {
    "Hello from Rust REST API!"
}