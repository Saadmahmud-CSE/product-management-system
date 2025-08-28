import React, { useState, useEffect } from 'react';
import '../styles/AddEditProduct.css';

function AddEditProduct({ product, onSave, onCancel }) {
  const [formData, setFormData] = useState({
    title: '',
    price: '',
    description: '',
    category: '',
    image: '',
    rating: {
      rate: '',
      count: '',
    },
  });

  useEffect(() => {
    if (product) {
      setFormData({
        title: product.title,
        price: product.price,
        description: product.description,
        category: product.category,
        image: product.image,
        rating: {
          rate: product.rating?.rate || '',
          count: product.rating?.count || '',
        },
      });
    }
  }, [product]);

  const handleChange = (e) => {
  const { name, value } = e.target;
  if (name === "rate" || name === "count") {
    setFormData((prev) => ({
      ...prev,
      rating: {
        ...prev.rating,
        [name]: name === "rate" ? parseFloat(value) || '' : parseInt(value) || '',
      }
    }));
  } else {
    setFormData((prev) => ({
      ...prev,
      [name]: name === 'price' ? parseFloat(value) || '' : value,
    }));
  }
};

  const handleSubmit = async (e) => {
    e.preventDefault();
    await onSave(formData);
    setFormData({
      title: '',
      price: '',
      description: '',
      category: '',
      image: '',
      rating: {
        rate: '',
        count: '',
      },
    });
  };

  return (
    <div className="add-edit-product">
      <h2>{product ? 'Edit Product' : 'Add Product'}</h2>
      <form onSubmit={handleSubmit}>
        <label>
          Title:
          <input
            type="text"
            name="title"
            value={formData.title}
            onChange={handleChange}
          />
        </label>
        <label>
          Price:
          <input
            type="number"
            name="price"
            value={formData.price}
            onChange={handleChange}
          />
        </label>
        <label>
          Description:
          <textarea
            name="description"
            value={formData.description}
            onChange={handleChange}
          />
        </label>
        <label>
          Category:
          <input
            type="string"
            name="category"
            value={formData.category}
            onChange={handleChange}
          />
        </label>
        <label>
          Image URL:
          <input
            type="text"
            name="image"
            value={formData.image}
            onChange={handleChange}
          />
        </label>
        <label>
          Rating (Rate):
          <input
            type="number"
            name="rate"
            step="0.1"
            min="0"
            max="5"
            value={formData.rating.rate}
            onChange={handleChange}
          />
        </label>
        <label>
          Rating (Count):
          <input
            type="number"
            name="count"
            min="0"
            value={formData.rating.count}
            onChange={handleChange}
          />
        </label>

        <div className="buttons">
          <button type="submit">{product ? 'Save Changes' : 'Add Product'}</button>
          <button type="button" onClick={onCancel}>
            Cancel
          </button>
        </div>
      </form>
    </div>
  );
}

export default AddEditProduct;
