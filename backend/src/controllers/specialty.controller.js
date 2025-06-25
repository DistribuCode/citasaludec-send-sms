const Specialty = require('../models/specialty.model');

exports.createSpecialty = async (req, res) => {
  try {
    const { name, description } = req.body;
    const specialty = new Specialty({ name, description });
    await specialty.save();
    res.status(201).json(specialty);
  } catch (error) {
    res.status(500).json({ error: error.message });
  }
};

exports.getAllSpecialties = async (req, res) => {
  try {
    const specialties = await Specialty.find();
    res.json(specialties);
  } catch (error) {
    res.status(500).json({ error: error.message });
  }
};
