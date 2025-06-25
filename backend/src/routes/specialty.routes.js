const express = require('express');
const router = express.Router();
const controller = require('../controllers/specialty.controller');

router.post('/', controller.createSpecialty);
router.get('/', controller.getAllSpecialties);

module.exports = router;
