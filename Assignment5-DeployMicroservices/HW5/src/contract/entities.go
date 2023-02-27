package contract

type GetFilesRequest struct {
	Unused     string
}

type GetFilesResponse struct {
        Filenames  []string
	Status     int
}

type GetFileRequest struct {
        Filename   string
}

type GetFileResponse struct {
        Filename   string
        Type       string
        Timestamp  string
        Length     string
        Data       string
	Status     int
}

type CreateFileRequest struct {
        Filename   string
        Type       string
        Timestamp  string
        Length     string
        Data       string
}

type CreateFileResponse struct {
	Status     int
}

type DeleteFileRequest struct {
        Filename   string
}

type DeleteFileResponse struct {
	Status     int
}
