const express = require('express');
const mongoose = require('mongoose');
const cors = require('cors');
const dotenv = require('dotenv');
const specialtyRoutes = require('./routes/specialty.routes');

dotenv.config();
const app = express();
app.use(cors());
app.use(express.json());

mongoose.connect(process.env.MONGO_URI)
  .then(() => console.log('✅ Connected to MongoDB'))
  .catch(err => console.error('❌ MongoDB connection error:', err));

app.use('/api/specialties', specialtyRoutes);

const PORT = process.env.PORT || 5010;
app.listen(PORT, () => {
  console.log(`🚀 Specialty Service running on port ${PORT}`);
});
