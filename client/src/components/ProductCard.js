import React from 'react';
import { Link } from 'react-router-dom';
import '../styles/ProductCard.css';

function ProductCard({ product, onEdit, onDelete }) {
  return (
    <div className="product-card">
      <div className="product-image">
        <img src={product.image} alt={product.title} />
      </div>
      <div className="product-details">
        <h3>{product.title}</h3>
        <p>Price: ${product.price}</p>
        <p>{product.description}</p>
        <p>{product.category}</p>
        {product.rating && (
          <p>
            {product.rating.rate} ⭐ ({product.rating.count} reviews)
          </p>
        )}
      </div>
      <div className="product-actions">
        <button onClick={() => onEdit(product.id)}>Edit</button>
        <button onClick={() => onDelete(product.id)}>Delete</button>
        <Link to={`/product/${product.id}`}>View Details</Link> {/* Add this link */}
      </div>
    </div>
  );
}

export default ProductCard;
