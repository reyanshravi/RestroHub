type Feedback struct {
    ID          uint      `gorm:"primaryKey"`
    Rating      int       `json:"rating"`
    Comment     string    `json:"comment" gorm:"type:text"`
    Language    string    `json:"language" gorm:"default:'en'"`
    StoreID     uint      `json:"store_id"`
    CreatedAt   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
    UpdatedAt   time.Time `json:"updated_at"`
}

// إضافة دعم التصدير للـ CSV
func (f *Feedback) ToCSVRecord() []string {
    return []string{
        strconv.Itoa(int(f.ID)),
        strconv.Itoa(f.Rating),
        f.Comment,
        f.Language,
        f.CreatedAt.Format(time.RFC3339),
    }
}